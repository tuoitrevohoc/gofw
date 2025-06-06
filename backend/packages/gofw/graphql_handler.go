package gofw

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

func GQLHandler(schema graphql.ExecutableSchema) *handler.Server {
	srv := handler.NewDefaultServer(schema)
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		logger := ContextLogger(ctx)
		statsd := ContextStatsd(ctx)

		statsd.Incr("graphql_error", []string{"path:" + err.Path.String()}, 1)
		logger.Error("GraphQL error", zap.Error(err))
		return err
	})

	return srv
}
