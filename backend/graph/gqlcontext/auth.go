package gqlcontext

import (
	"context"
	"net/http"

	"github.com/go-pkgz/auth/token"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user_sg"}

type contextKey struct {
	name string
}

// Auth Middleware decodes the share session cookie and packs the session into context
func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := token.GetUserInfo(r)

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func UserFromContext(ctx context.Context) token.User {
	raw, _ := ctx.Value(userCtxKey).(token.User)
	return raw
}
