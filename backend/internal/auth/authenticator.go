package auth

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/user"
	"github.com/tuoitrevohoc/gofw/backend/packages/cache"
	"github.com/tuoitrevohoc/gofw/backend/packages/gofw"
	"go.uber.org/zap"
)

const (
	sessionCacheKey = "auth_session_%s"
	sessionCacheTTL = 10 * time.Minute
)

type Authenticator struct {
	ent        *ent.Client
	webauthn   *webauthn.WebAuthn
	timedCache cache.TimedCache
}

func NewPasskeyAuthenticator(manager *gofw.ServiceManager,
	ent *ent.Client,
	domain string,
	origin string,
	timedCache cache.TimedCache,
) *Authenticator {

	if len(origin) == 0 {
		origin = fmt.Sprintf("https://%s", domain)
	}

	manager.Logger().
		Info("Initializing passkey authenticator",
			zap.String("domain", domain),
			zap.String("origin", origin),
		)

	webauthn, err := webauthn.New(&webauthn.Config{
		RPID:          domain,
		RPDisplayName: "Gofw",
		RPOrigins:     []string{origin},
		Timeouts: webauthn.TimeoutsConfig{
			Registration: webauthn.TimeoutConfig{
				Enforce:    true,
				Timeout:    10 * time.Minute,
				TimeoutUVD: 10 * time.Minute,
			},
			Login: webauthn.TimeoutConfig{
				Enforce:    true,
				Timeout:    10 * time.Minute,
				TimeoutUVD: 10 * time.Minute,
			},
		},
	})

	if err != nil {
		manager.Logger().Fatal("Passkey authenticator initialized", zap.Error(err))
	}

	return &Authenticator{
		ent:        ent,
		webauthn:   webauthn,
		timedCache: timedCache,
	}
}

func (a *Authenticator) checkUserExists(ctx context.Context, email string) (bool, error) {
	logger := gofw.ContextLogger(ctx)
	logger.Info("checkUserExists")

	exists, err := a.ent.User.Query().Where(user.Email(email)).Exist(ctx)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (a *Authenticator) BeginRegistration(ctx context.Context, email string) (*protocol.CredentialCreation, error) {
	logger := gofw.ContextLogger(ctx)
	logger.Info("BeginRegistration")

	userExists, err := a.checkUserExists(ctx, email)
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, errors.New("user already exists")
	}

	creation, session, err := a.webauthn.BeginRegistration(NewWebAuthnUser(email))
	if err != nil {
		logger.Error("error beginning registration", zap.Error(err))
		return nil, err
	}

	session = encodeSession(session)
	timedCacheKey := fmt.Sprintf(sessionCacheKey, session.Challenge)
	err = a.timedCache.Set(ctx, timedCacheKey, session, sessionCacheTTL)

	if err != nil {
		logger.Error("error setting timed cache", zap.Error(err))
		return nil, err
	}

	logger.Info("BeginRegistration success", zap.String("timedCacheKey", timedCacheKey))
	return creation, nil
}

func (a *Authenticator) FinishRegistration(ctx context.Context, email string, response *protocol.ParsedCredentialCreationData) (*ent.User, error) {
	logger := gofw.ContextLogger(ctx)
	logger.Info("FinishRegistration")

	userExists, err := a.checkUserExists(ctx, email)
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, errors.New("user already exists")
	}

	timedCacheKey := fmt.Sprintf(sessionCacheKey, response.Response.CollectedClientData.Challenge)
	var authnSession *webauthn.SessionData

	err = a.timedCache.Get(ctx, timedCacheKey, &authnSession)
	logger.Info("FinishRegistration", zap.String("timedCacheKey", timedCacheKey), zap.Any("authnSession", authnSession))

	if err != nil {
		logger.Error("error getting timed cache", zap.Error(err))
		return nil, err
	}

	if string(authnSession.UserID) != email {
		logger.Error("user ID mismatch", zap.String("expected", email), zap.String("actual", string(authnSession.UserID)))
		return nil, errors.New("user ID mismatch")
	}

	credential, err := a.webauthn.CreateCredential(NewWebAuthnUser(email), *authnSession, response)
	if err != nil {
		logger.Error("Error finishing registration", zap.Error(err))
		return nil, err
	}

	entCredential, err := convertToCredential(credential)
	if err != nil {
		logger.Error("Error converting to ent credential", zap.Error(err))
		return nil, err
	}

	// create user
	user, err := a.ent.User.Create().SetEmail(email).SetFinishedRegistration(true).Save(ctx)
	if err != nil {
		logger.Error("Error creating user", zap.Error(err))
		return nil, err
	}

	_, err = a.ent.Credential.Create().SetData(entCredential.Data).SetPublicKey(entCredential.PublicKey).SetUser(user).Save(ctx)

	if err != nil {
		logger.Error("Error saving credential", zap.Error(err))
		return nil, err
	}

	return user, nil
}

func (a *Authenticator) BeginLogin(ctx context.Context) (*protocol.CredentialAssertion, error) {
	logger := gofw.ContextLogger(ctx)
	logger.Info("BeginLogin")

	var request *protocol.CredentialAssertion
	var session *webauthn.SessionData
	var err error

	request, session, err = a.webauthn.BeginDiscoverableLogin()

	if err != nil {
		logger.Error("error beginning login", zap.Error(err))
		return nil, err
	}

	session = encodeSession(session)
	timedCacheKey := fmt.Sprintf(sessionCacheKey, session.Challenge)
	err = a.timedCache.Set(ctx, timedCacheKey, session, sessionCacheTTL)

	if err != nil {
		logger.Error("error setting timed cache", zap.Error(err))
		return nil, err
	}

	logger.Info("BeginLogin success", zap.String("timedCacheKey", timedCacheKey))
	return request, nil
}

func (a *Authenticator) FinishLogin(ctx context.Context, response *protocol.ParsedCredentialAssertionData) (*ent.User, error) {
	logger := gofw.ContextLogger(ctx)
	logger.Info("FinishLogin", zap.String("userHandler", string(response.Response.UserHandle)))

	timedCacheKey := fmt.Sprintf(sessionCacheKey, response.Response.CollectedClientData.Challenge)
	var authnSession webauthn.SessionData

	userHandler, err := base64.StdEncoding.DecodeString(string(response.Response.UserHandle))
	if err != nil {
		logger.Error("error decoding user handle", zap.Error(err))
		userHandler = response.Response.UserHandle
	}

	err = a.timedCache.Get(ctx, timedCacheKey, &authnSession)
	if err != nil {
		logger.Error("error getting timed cache", zap.Error(err))
		return nil, err
	}

	user, err := a.ent.User.Query().Where(user.Email(string(userHandler))).WithCredentials().Only(ctx)
	if err != nil {
		return nil, err
	}

	webUser, err := NewWebAuthnUserFromUser(ctx, user)
	if err != nil {
		return nil, err
	}

	authnSession.UserID = response.Response.UserHandle

	_, err = a.webauthn.ValidateLogin(webUser, authnSession, response)

	if err != nil {
		logger.Error("error validating login", zap.Error(err))
		return nil, err
	}

	return user, nil
}
