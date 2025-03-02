package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
	gql "github.com/tuoitrevohoc/gofw/backend/gen/go/graphql"
)

// NewHandler returns a new GraphQL handler.
func NewHandler(client *ent.Client) *handler.Server {
	return handler.NewDefaultServer(gql.NewExecutableSchema(gql.Config{Resolvers: &Resolver{client: client}}))
}
