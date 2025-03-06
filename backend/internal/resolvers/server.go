package resolvers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/tuoitrevohoc/gofw/backend/gen/go/ent"
	gql "github.com/tuoitrevohoc/gofw/backend/gen/go/graphql"
	"github.com/tuoitrevohoc/gofw/backend/packages/gofw"
)

// NewHandler returns a new GraphQL handler.
func NewHandler(client *ent.Client) *handler.Server {
	return gofw.GQLHandler(gql.NewExecutableSchema(gql.Config{Resolvers: &Resolver{client: client}}))
}
