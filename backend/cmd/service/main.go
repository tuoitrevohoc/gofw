package main

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tuoitrevohoc/gofw/backend/internal/auth"
	"github.com/tuoitrevohoc/gofw/backend/internal/config"
	"github.com/tuoitrevohoc/gofw/backend/internal/db"
	"github.com/tuoitrevohoc/gofw/backend/internal/resolvers"
	"github.com/tuoitrevohoc/gofw/backend/packages/gofw"
	"go.uber.org/zap"
)

func main() {
	manager := gofw.NewServiceManager()
	logger := manager.Logger()

	cfg, err := config.LoadAppConfig()

	if err != nil {
		logger.Fatal("Failed to get app config", zap.Error(err))
	}

	entClient, err := db.NewEntClient(cfg.DbUrl)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	// migrate database
	if err := entClient.Schema.Create(context.Background()); err != nil {
		logger.Fatal("Failed to create schema", zap.Error(err))
	}

	authenticator := auth.NewPasskeyAuthenticator(manager, entClient, cfg.Domain, cfg.Origin)
	authenticationService := auth.NewAuthenticationService(entClient, []byte(cfg.JwtSecret))
	graphqlHandler := resolvers.NewHandler(entClient, authenticator, authenticationService)
	server := gofw.NewHttpServer(manager, cfg.Port)

	server.Use(authenticationService.AuthMiddleware)

	server.AddHandler("/graphql", graphqlHandler)
	server.AddHandler("/graphql/playground", playground.Handler("GraphQL playground", "/graphql"))
	server.AddHandler(auth.RefreshTokenPath, http.HandlerFunc(authenticationService.RefreshTokenEndpoint))

	manager.AddService(server)

	err = manager.Run()
	if err != nil {
		logger.Error("Failed to run service manager", zap.Error(err))
	}
}
