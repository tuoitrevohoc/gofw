package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
)

type WebAuthnUser struct {
	email string
	creds []webauthn.Credential
}

func NewWebAuthnUser(email string) *WebAuthnUser {
	return &WebAuthnUser{
		email: email,
		creds: []webauthn.Credential{},
	}
}

func NewWebAuthnUserFromUser(ctx context.Context, user *ent.User) (*WebAuthnUser, error) {
	creds, err := user.Credentials(ctx)
	if err != nil {
		return nil, err
	}

	webCreds := make([]webauthn.Credential, len(creds))
	for i, cred := range creds {
		webCred, err := convertToWebAuthnCredential(cred)
		if err != nil {
			return nil, err
		}
		webCreds[i] = *webCred
	}

	encodedEmail := base64.RawStdEncoding.EncodeToString([]byte(user.Email))

	return &WebAuthnUser{
		email: encodedEmail,
		creds: webCreds,
	}, nil
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
	return u.creds
}

func encodeSession(session *webauthn.SessionData) *webauthn.SessionData {
	challenge := base64.RawStdEncoding.EncodeToString([]byte(session.Challenge))
	session.Challenge = challenge
	return session
}

func convertToCredential(webauthnCredential *webauthn.Credential) (*ent.Credential, error) {
	// convert public key to base64
	publicKey := base64.RawStdEncoding.EncodeToString(webauthnCredential.PublicKey)

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
