package gofw

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

// Middleware is a function that wraps an http.Handler and returns a new http.Handler
type Middleware func(http.Handler) http.Handler

type HttpServer struct {
	port       string
	mux        *http.ServeMux
	manager    *ServiceManager
	middleware []Middleware
}

func NewHttpServer(manager *ServiceManager, port string) *HttpServer {
	mux := http.NewServeMux()
	defaultMiddlewares := []Middleware{
		ServiceManagerMiddleware(manager),
	}

	server := &HttpServer{
		port:       port,
		mux:        mux,
		manager:    manager,
		middleware: defaultMiddlewares,
	}

	return server
}

// Use adds middleware to the server. Middleware will be executed in the order they are added.
func (s *HttpServer) Use(middleware ...Middleware) {
	s.middleware = append(s.middleware, middleware...)
}

func (s *HttpServer) AddHandler(path string, handler http.Handler) {
	s.mux.HandleFunc(path, handler.ServeHTTP)
}

// chainMiddleware chains all registered middleware with the final handler
func (s *HttpServer) chainMiddleware(finalHandler http.Handler) http.Handler {
	// Apply middleware in reverse order so that the first middleware in the slice
	// is the outermost (first to handle the request)
	handler := finalHandler
	for i := len(s.middleware) - 1; i >= 0; i-- {
		handler = s.middleware[i](handler)
	}
	return handler
}

func (s *HttpServer) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":" + s.port,
		Handler: s.chainMiddleware(s.mux), // Apply middleware chain to the mux
	}

	// Channel to capture server errors
	errChan := make(chan error, 1)

	// Start server in a goroutine
	go func() {
		s.manager.Logger().Info("Starting HTTP server on port ", zap.String("port", s.port))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	// Wait for context cancellation or server error
	select {
	case <-ctx.Done():
		// Graceful shutdown
		return server.Shutdown(context.Background())
	case err := <-errChan:
		return err
	}
}
