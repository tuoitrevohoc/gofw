package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
)

type WebAuthnUser struct {
	ent.User
}

func NewWebAuthnUser(user *ent.User) *WebAuthnUser {
	return &WebAuthnUser{
		User: *user,
	}
}

func (u *WebAuthnUser) WebAuthnID() []byte {
	return []byte(fmt.Sprintf("%d", u.ID))
}

func (u *WebAuthnUser) WebAuthnName() string {
	return u.Email
}

func (u *WebAuthnUser) WebAuthnDisplayName() string {
	return u.Email
}

func (u *WebAuthnUser) WebAuthnCredentials() []webauthn.Credential {
	return []webauthn.Credential{}
}

func convertToAuthSession(session *webauthn.SessionData) (*ent.AuthSession, error) {
	data, err := json.Marshal(session)
	if err != nil {
		return nil, err
	}

	return &ent.AuthSession{
		Data: string(data),
	}, nil
}

func convertToCredential(webauthnCredential *webauthn.Credential) (*ent.Credential, error) {
	// convert public key to base64
	publicKey := base64.StdEncoding.EncodeToString(webauthnCredential.PublicKey)

	// convert data to base64
	data, err := json.Marshal(webauthnCredential)
	if err != nil {
		return nil, err
	}

	return &ent.Credential{
		PublicKey: publicKey,
		Data:      string(data),
	}, nil
}

func convertToWebAuthnCredential(credential *ent.Credential) (*webauthn.Credential, error) {
	var result webauthn.Credential

	if err := json.Unmarshal([]byte(credential.Data), &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func convertToSessionData(session *ent.AuthSession) (*webauthn.SessionData, error) {
	var result webauthn.SessionData

	if err := json.Unmarshal([]byte(session.Data), &result); err != nil {
		return nil, err
	}

	// decode challenge
	result.Challenge = base64.URLEncoding.EncodeToString([]byte(result.Challenge))
	result.Challenge = strings.TrimRight(result.Challenge, "=")
	return &result, nil
}
