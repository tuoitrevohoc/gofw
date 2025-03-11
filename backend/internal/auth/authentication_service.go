package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/refreshtoken"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent/user"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/graphql/model"
	"github.com/tuoitrevohoc/gofw/backend/internal/scalars"
	"github.com/tuoitrevohoc/gofw/backend/packages/gofw"
)

const (
	refreshTokenCookieName   = "refresh_token"
	refreshTokenCookieMaxAge = 3600 * 24 * 30 * time.Second // 30 days
	accessTokenMaxAge        = 15 * 60 * time.Second        // 15 minutes
	RefreshTokenPath         = "/api/refresh-token"
)

type AuthenticationService struct {
	ent       *ent.Client
	jwtSecret []byte
}

type AccessToken struct {
	AccessToken string `json:"accessToken"`
	Expiry      int64  `json:"expiry"`
}

func NewAuthenticationService(ent *ent.Client, jwtSecret []byte) *AuthenticationService {
	return &AuthenticationService{ent: ent, jwtSecret: jwtSecret}
}

func (s *AuthenticationService) RefreshTokenEndpoint(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	token, err := s.RefreshToken(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response, err := json.Marshal(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *AuthenticationService) Logout(ctx context.Context) error {
	defer func() {
		gofw.ContextStatsd(ctx).Incr("auth_logout", []string{}, 1)
	}()

	writer := gofw.ContextResponseWriter(ctx)
	viewer := ViewerFromContext(ctx)

	if viewer == nil {
		return errors.New("viewer not found")
	}

	_, err := s.ent.RefreshToken.Update().
		Where(refreshtoken.ID(*viewer.SessionID)).
		SetIsActive(false).
		Save(ctx)

	if err != nil {
		return err
	}

	http.SetCookie(*writer, &http.Cookie{
		Name:     refreshTokenCookieName,
		Value:    "",
		MaxAge:   -1,
		Path:     RefreshTokenPath,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}

func (s *AuthenticationService) setRefreshTokenCookie(ctx context.Context, refreshToken *ent.RefreshToken) error {
	writer := gofw.ContextResponseWriter(ctx)
	cookie := &http.Cookie{
		Name:     refreshTokenCookieName,
		Value:    refreshToken.Token,
		MaxAge:   int(refreshTokenCookieMaxAge.Seconds()),
		HttpOnly: true,
		Path:     RefreshTokenPath,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(*writer, cookie)

	return nil
}

func (s *AuthenticationService) Login(ctx context.Context, user *ent.User) (*AccessToken, error) {
	defer func() {
		gofw.ContextStatsd(ctx).Incr("auth_login", []string{}, 1)
	}()

	// generate access token
	request := gofw.ContextRequest(ctx)
	ipAddress := request.RemoteAddr
	userAgent := request.UserAgent()

	refreshToken, err := s.ent.RefreshToken.Create().
		SetToken(uuid.New().String()).
		SetExpireAt(time.Now().Add(refreshTokenCookieMaxAge)).
		SetUserID(user.ID).
		SetIPAddress(ipAddress).
		SetUserAgent(userAgent).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	err = s.setRefreshTokenCookie(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	return s.GenerateAccessToken(ctx, refreshToken, user)
}

func (s *AuthenticationService) RefreshToken(ctx context.Context) (*AccessToken, error) {
	defer func() {
		gofw.ContextStatsd(ctx).Incr("auth_refresh_token", []string{}, 1)
	}()

	request := gofw.ContextRequest(ctx)

	tokenCookie, err := request.Cookie(refreshTokenCookieName)

	if err != nil || tokenCookie == nil {
		return nil, errors.New("refresh token not found")
	}

	refreshToken, err := s.ent.RefreshToken.Query().
		Where(refreshtoken.Token(tokenCookie.Value), refreshtoken.IsActive(true)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	if refreshToken == nil {
		return nil, errors.New("refresh token not found")
	}

	if refreshToken.ExpireAt.Before(time.Now()) {
		return nil, errors.New("refresh token expired")
	}

	user, err := refreshToken.QueryUser().Only(ctx)
	if err != nil {
		return nil, err
	}

	err = s.setRefreshTokenCookie(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	return s.GenerateAccessToken(ctx, refreshToken, user)
}

func (s *AuthenticationService) GenerateAccessToken(ctx context.Context, refreshToken *ent.RefreshToken, usr *ent.User) (*AccessToken, error) {
	expiry := time.Now().Add(accessTokenMaxAge)

	claims := jwt.MapClaims{
		"sub":        scalars.NewGUID(user.Table, usr.ID).String(),
		"exp":        expiry.Unix(),
		"iat":        time.Now().Unix(),
		"refresh_id": refreshToken.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)

	if err != nil {
		return nil, err
	}

	return &AccessToken{AccessToken: tokenString, Expiry: expiry.Unix()}, nil
}

func (s *AuthenticationService) VerifyAccessToken(ctx context.Context, token string) (*model.Viewer, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return s.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	refreshIDNumber, ok := claims["refresh_id"].(float64)
	if !ok {
		return nil, errors.New("refresh id not found")
	}

	refreshID := int(refreshIDNumber)

	return &model.Viewer{
		UserID:          scalars.ParseGUID(claims["sub"].(string)),
		IsAuthenticated: true,
		SessionID:       &refreshID,
	}, nil
}
