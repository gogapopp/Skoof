package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gogapopp/Skoof/internal/libs/jwt"
)

type (
	CtxKeyUserRole string
	CtxKeyUserID   int
)

const (
	UserRoleKey CtxKeyUserRole = ""
	UserIDKey   CtxKeyUserID   = 0
)

// AuthMiddleware checks the user cookie for the JWT token.
func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("ssid")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				http.Redirect(w, r, "/signin?redirected=true", http.StatusTemporaryRedirect)
				return
			default:
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
		}

		jwtToken := cookie.Value

		userID, role, err := jwt.ParseJWTToken(jwtToken)
		if err != nil {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserRoleKey, fmt.Sprint(role))
		ctx = context.WithValue(ctx, UserIDKey, fmt.Sprint(userID))

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
