package graphql

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
)

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

func parseGUID(_ context.Context, guid string) (int, string, error) {
	parts := strings.Split(guid, "/")

	if len(parts) != 2 {
		return 0, "", fmt.Errorf("unexpected id format. expect Type/ID, but got: %s", guid)
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, "", fmt.Errorf("invalid id: %s", parts[1])
	}

	return id, parts[0], nil
}
