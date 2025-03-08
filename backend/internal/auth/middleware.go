package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/tuoitrevohoc/gofw/backend/gen/go/graphql/model"
)

type viewerKey struct{}

func WithViewer(ctx context.Context, viewer *model.Viewer) context.Context {
	return context.WithValue(ctx, viewerKey{}, viewer)
}

func ViewerFromContext(ctx context.Context) *model.Viewer {
	loggedInViewer := ctx.Value(viewerKey{})

	if loggedInViewer == nil {
		return &model.Viewer{
			IsAuthenticated: false,
		}
	}

	return loggedInViewer.(*model.Viewer)
}

func (s *AuthenticationService) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}

		viewer, err := s.VerifyAccessToken(r.Context(), token)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		ctx := WithViewer(r.Context(), viewer)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
