package handler

import (
	"context"
	"net/http"

	"github.com/gogapopp/Skoof/components"
	"github.com/gogapopp/Skoof/model"
)

type Service interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, emailOrUsername, password string) (model.User, error)
}

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
