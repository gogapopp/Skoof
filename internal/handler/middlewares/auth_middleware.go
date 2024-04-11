package middlewares

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gogapopp/Skoof/internal/lib/jwt"
)

type CtxKeyUserID int

const UserIDKey CtxKeyUserID = 0

func AuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("ssid")
		if err != nil {
			switch {
			case errors.Is(err, http.ErrNoCookie):
				http.Error(w, "cookie not found", http.StatusBadRequest)
			default:
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}
		}

		jwtToken := cookie.Value

		userID, err := jwt.ParseJWTToken(jwtToken)
		if err != nil {
			http.Error(w, "invalid authorization token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, fmt.Sprint(userID))

		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
