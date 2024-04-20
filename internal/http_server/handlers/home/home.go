package home

import (
	"net/http"

	"github.com/gogapopp/Skoof/components"
	"github.com/gogapopp/Skoof/internal/libs/render"
	"go.uber.org/zap"
)

func HomePage(logger *zap.SugaredLogger) http.HandlerFunc {
	const op = "handler.home_page.HomePage"
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_ = ctx
		switch r.Method {
		case http.MethodGet:
			if err := render.Render(r.Context(), w, components.HomeBase(components.Home("For skoofs from Skoof"))); err != nil {
				logger.Errorf("%s: %w", op, err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}
