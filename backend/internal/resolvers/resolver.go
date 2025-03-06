package resolvers

import (
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
	"github.com/tuoitrevohoc/gofw/backend/internal/auth"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct {
	client        *ent.Client
	authenticator *auth.Authenticator
}

// NewResolver creates a new resolver.
func NewResolver(client *ent.Client, authenticator *auth.Authenticator) *Resolver {
	return &Resolver{client: client, authenticator: authenticator}
}
