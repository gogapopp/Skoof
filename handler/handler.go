package handler

import (
	"net/http"

	"github.com/gogapopp/Skoof/components"
	"github.com/gogapopp/Skoof/handler/middlewares"
)

func HomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if err := render(r.Context(), w, components.Home("Skoof")); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func SkoofPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			v, ok := r.Context().Value(middlewares.UserIDKey).(string)
			if !ok {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			_ = v
			// TODO:
			if err := render(r.Context(), w, components.Home("Skoof")); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}
