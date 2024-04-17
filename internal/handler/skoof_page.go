package handler

import (
	"net/http"

	"github.com/gogapopp/Skoof/components/skoof_pages"
	"github.com/gogapopp/Skoof/internal/handler/middlewares"
	"go.uber.org/zap"
)

func SkoofPage(logger *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_ = ctx
		switch r.Method {
		case http.MethodGet:
			v, ok := r.Context().Value(middlewares.UserIDKey).(string)
			if !ok {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			_ = v
			// TODO:
			if err := render(r.Context(), w, skoof_pages.SkoofBase(skoof_pages.Skoof())); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}
