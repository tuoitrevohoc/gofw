package gofw

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

type HealthServer struct {
	server         *http.Server
	serviceManager *ServiceManager
}

func NewHealthServer(serviceManager *ServiceManager, port string) *HealthServer {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		serviceManager.Statsd().Incr("health_server.requests", nil, 1)
	})

	return &HealthServer{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: mux,
		},
		serviceManager: serviceManager,
	}
}

func (s *HealthServer) Run(ctx context.Context) error {
	s.serviceManager.Logger().Info("Starting health server", zap.String("address", s.server.Addr))
	return s.server.ListenAndServe()
}
