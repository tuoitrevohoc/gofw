package gofw

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type HealthServer struct {
	port             string
	router           chi.Router
	serviceManager *ServiceManager
}

func NewHealthServer(serviceManager *ServiceManager, port string) *HealthServer {
	router := chi.NewRouter()
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		serviceManager.Statsd().Incr("health_server.requests", nil, 1)
	})

	return &HealthServer{
		router:           router,
		serviceManager: serviceManager,
		port:             port,
	}
}

func (s *HealthServer) Run(ctx context.Context) error {
	return http.ListenAndServe(":"+s.port, s.router)
}
