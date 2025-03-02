package graphql

import "github.com/tuoitrevohoc/gofw/backend/gen/go/ent"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct {
	client *ent.Client
}

// NewResolver creates a new resolver.
func NewResolver(client *ent.Client) *Resolver {
	return &Resolver{client: client}
}
