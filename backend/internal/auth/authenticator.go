package auth

import (
	"context"
	"errors"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/authsession"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/user"
	"github.com/tuoitrevohoc/gofw/backend/packages/gofw"
	"go.uber.org/zap"
)

type Authenticator struct {
	ent      *ent.Client
	webauthn *webauthn.WebAuthn
}

func NewPasskeyAuthenticator(manager *gofw.ServiceManager, ent *ent.Client, domain string, origin string) *Authenticator {
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
		ent:      ent,
		webauthn: webauthn,
	}
}

func (a *Authenticator) findUserToRegister(ctx context.Context, email string) (*ent.User, error) {
	logger := gofw.ContextLogger(ctx)
	logger.Info("findUserToRegister")

	user, err := a.ent.User.Query().Where(user.Email(email)).Only(ctx)

	// if not found, create a new user
	if err != nil && ent.IsNotFound(err) {
		user, err = a.ent.User.Create().SetEmail(email).SetFinishedRegistration(false).Save(ctx)
		if err != nil {
			logger.Error("error creating user", zap.Error(err))
			return nil, err
		}
	} else if err != nil {
		logger.Error("error querying user", zap.Error(err))
		return nil, err
	}

	if user.FinishedRegistration {
		logger.Error("user already finished registration")
		return nil, errors.New("user already exists")
	}

	return user, nil
}

func (a *Authenticator) createOrUpdateSession(ctx context.Context, user *ent.User, data string) error {
	logger := gofw.ContextLogger(ctx)
	logger.Info("createOrUpdateSession")

	existingSession, err := a.ent.AuthSession.Query().
		Where(authsession.UserID(user.ID)).
		Only(ctx)

	if err != nil && ent.IsNotFound(err) {
		// Create new session if not found
		_, err = a.ent.AuthSession.Create().
			SetData(data).
			SetUser(user).
			Save(ctx)
	} else if err == nil {
		// Update existing session
		_, err = existingSession.Update().
			SetData(data).
			Save(ctx)
	}

	if err != nil {
		logger.Error("Error creating/updating session", zap.Error(err))
		return err
	}

	return nil
}

func (a *Authenticator) BeginRegistration(ctx context.Context, email string) (*protocol.CredentialCreation, error) {
	logger := gofw.ContextLogger(ctx)
	logger.Info("BeginRegistration")

	user, err := a.findUserToRegister(ctx, email)
	if err != nil {
		return nil, err
	}

	creation, session, err := a.webauthn.BeginRegistration(NewWebAuthnUser(user))

	if err != nil {
		logger.Error("Error beginning registration", zap.Error(err))
		return nil, err
	}

	authSession, err := convertToAuthSession(session)
	logger.Info("authSession", zap.Any("challenge", session.Challenge), zap.Any("response", creation.Response.Challenge))

	if err != nil {
		logger.Error("Error converting to auth session", zap.Error(err))
		return nil, err
	}

	err = a.createOrUpdateSession(ctx, user, authSession.Data)
	if err != nil {
		return nil, err
	}

	return creation, nil
}

func (a *Authenticator) FinishRegistration(ctx context.Context, email string, response *protocol.ParsedCredentialCreationData) (*ent.User, error) {
	logger := gofw.ContextLogger(ctx)
	logger.Info("FinishRegistration")

	user, err := a.ent.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		logger.Error("Error querying user", zap.Error(err))
		return nil, err
	}

	if user.FinishedRegistration {
		logger.Error("User already finished registration")
		return nil, errors.New("user already exists")
	}

	session, err := a.ent.AuthSession.Query().
		Where(authsession.UserID(user.ID)).
		Only(ctx)

	if err != nil {
		logger.Error("Error querying auth session", zap.Error(err), zap.String("challenge", response.Response.CollectedClientData.Challenge))
		return nil, err
	}

	authnSession, err := convertToSessionData(session)

	if err != nil {
		logger.Error("Error converting to authn session", zap.Error(err))
		return nil, err
	}

	credential, err := a.webauthn.CreateCredential(NewWebAuthnUser(user), *authnSession, response)

	if err != nil {
		logger.Error("Error finishing registration", zap.Error(err))
		return nil, err
	}

	entCredential, err := convertToCredential(credential)

	if err != nil {
		logger.Error("Error converting to ent credential", zap.Error(err))
		return nil, err
	}

	_, err = a.ent.Credential.Create().SetData(entCredential.Data).SetPublicKey(entCredential.PublicKey).SetUser(user).Save(ctx)

	if err != nil {
		logger.Error("Error saving credential", zap.Error(err))
		return nil, err
	}

	user.FinishedRegistration = true
	_, err = a.ent.User.UpdateOne(user).SetFinishedRegistration(true).Save(ctx)
	if err != nil {
		logger.Error("Error updating user", zap.Error(err))
		return nil, err
	}

	return user, nil
}
