// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
	"github.com/tuoitrevohoc/gofw/backend/packages/scalars"
)

type AccessToken struct {
	AccessToken string  `json:"accessToken"`
	Expiry      int     `json:"expiry"`
	Viewer      *Viewer `json:"viewer"`
}

type AuthnLoginResponse struct {
	CredentialRequest string `json:"credentialRequest"`
}

type AuthnRegistrationResponse struct {
	CredentialCreation string `json:"credentialCreation"`
}

type SaveRestaurantInput struct {
	ID      *scalars.GUID `json:"id,omitempty"`
	Name    string        `json:"name"`
	Address string        `json:"address"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Viewer struct {
	Profile         *ent.User     `json:"profile"`
	UserID          *scalars.GUID `json:"userId,omitempty"`
	IsAuthenticated bool          `json:"isAuthenticated"`
	SessionID       *int          `json:"sessionId,omitempty"`
}
