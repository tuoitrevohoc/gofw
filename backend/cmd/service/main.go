package main

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tuoitrevohoc/gofw/backend/internal/config"
	"github.com/tuoitrevohoc/gofw/backend/internal/db"
	"github.com/tuoitrevohoc/gofw/backend/internal/graphql"
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

	entClient, err := db.NewEntClient(cfg.Db.Url)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}

	graphqlHandler := graphql.NewHandler(entClient)

	server := gofw.NewHttpServer(manager, cfg.Port)
	server.AddHandler("/graphql", graphqlHandler)
	server.AddHandler("/graphql/playground", playground.Handler("GraphQL playground", "/graphql"))

	manager.AddService(server)

	err = manager.Run()
	if err != nil {
		logger.Error("Failed to run service manager", zap.Error(err))
	}
}
