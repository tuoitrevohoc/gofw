package resolvers

import (
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
	"github.com/tuoitrevohoc/gofw/backend/internal/auth"
	"github.com/tuoitrevohoc/gofw/backend/packages/cache"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct {
	client                *ent.Client
	authenticator         *auth.Authenticator
	authenticationService *auth.AuthenticationService
	timedCache            cache.TimedCache
}

// NewResolver creates a new resolver.
func NewResolver(client *ent.Client, authenticator *auth.Authenticator, authenticationService *auth.AuthenticationService, timedCache cache.TimedCache) *Resolver {
	return &Resolver{
		client:                client,
		authenticator:         authenticator,
		authenticationService: authenticationService,
		timedCache:            timedCache,
	}
}
