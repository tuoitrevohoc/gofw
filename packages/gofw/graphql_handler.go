package gofw

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

func GQLHandler(schema graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(schema)
	srv.AddTransport(transport.SSE{
		KeepAlivePingInterval: 10 * time.Second,
	})

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})

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
