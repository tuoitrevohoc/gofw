package gofw

import (
	"context"
	"net/http"

	"github.com/DataDog/datadog-go/v5/statsd"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type loggerKey struct{}
type requestIDKey struct{}
type statsdKey struct{}

const RequestIDHeader = "X-Request-ID"

// FromContext retrieves the logger from the context
func FromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(loggerKey{}).(*zap.Logger); ok {
		return logger
	}
	return zap.NewNop()
}

// WithLogger adds a logger to the context
func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func WithStatsd(ctx context.Context, statsd *statsd.Client) context.Context {
	return context.WithValue(ctx, statsdKey{}, statsd)
}

// RequestIDFromContext retrieves the request ID from the context
func RequestIDFromContext(ctx context.Context) string {
	if requestID, ok := ctx.Value(requestIDKey{}).(string); ok {
		return requestID
	}
	return ""
}

// WithRequestID adds a request ID to the context
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, requestID)
}

// ServiceManagerMiddleware creates a new middleware that adds the logger to the request context
func ServiceManagerMiddleware(manager *ServiceManager) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get request ID from header or generate a new one
			requestID := r.Header.Get(RequestIDHeader)
			if requestID == "" {
				requestID = uuid.New().String()
			}

			// Add request ID to response header
			w.Header().Set(RequestIDHeader, requestID)

			// Create a logger with request ID
			requestLogger := manager.Logger().With(zap.String("request_id", requestID))

			// Add both logger and request ID to context
			ctx := WithLogger(r.Context(), requestLogger)
			ctx = WithStatsd(ctx, manager.Statsd())
			ctx = WithRequestID(ctx, requestID)

			// Add panic recovery
			defer func() {
				if err := recover(); err != nil {
					requestLogger.Error("panic occurred during request handling",
						zap.Any("error", err),
						zap.String("request_id", requestID),
						zap.String("path", r.URL.Path),
						zap.String("method", r.Method),
					)
					w.WriteHeader(http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
