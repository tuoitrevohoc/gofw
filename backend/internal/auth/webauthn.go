package auth

import (
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
)

type WebAuthnUser struct {
	email string
}

func NewWebAuthnUser(email string) *WebAuthnUser {
	return &WebAuthnUser{
		email: email,
	}
}

func (u *WebAuthnUser) WebAuthnID() []byte {
	return []byte(u.email)
}

func (u *WebAuthnUser) WebAuthnName() string {
	return u.email
}

func (u *WebAuthnUser) WebAuthnDisplayName() string {
	return u.email
}

func (u *WebAuthnUser) WebAuthnCredentials() []webauthn.Credential {
	return []webauthn.Credential{}
}

func encodeSession(session *webauthn.SessionData) *webauthn.SessionData {
	challenge := base64.StdEncoding.EncodeToString([]byte(session.Challenge))
	challenge = strings.TrimRight(challenge, "=")
	session.Challenge = challenge
	return session
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
