package httpserver

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gogapopp/Skoof/internal/config"
	"github.com/gogapopp/Skoof/internal/http_server/handlers/auth"
	"github.com/gogapopp/Skoof/internal/http_server/handlers/home"
	"github.com/gogapopp/Skoof/internal/http_server/handlers/skoof"
	"github.com/gogapopp/Skoof/internal/http_server/middlewares"
	"go.uber.org/zap"
)

const readHeaderTimeout = 10 * time.Second

func initHandlers(r *chi.Mux, logger *zap.SugaredLogger, authService auth.AuthService) {
	r.Get("/", home.HomePage(logger))

	r.Get("/signin", auth.SignInPage(logger, authService))
	r.Post("/signin", auth.SignInPage(logger, authService))

	r.Get("/signup", auth.SignUpPage(logger, authService))
	r.Post("/signup", auth.SignUpPage(logger, authService))

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)

		r.Get("/skoof", skoof.SkoofPage(logger))
		r.Post("/skoof", skoof.SkoofPage(logger))

		r.Get("/skoof/community", skoof.SkoofPage(logger))
		r.Post("/skoof/community", skoof.SkoofPage(logger))

		r.Get("/skoof/community/{id}", skoof.SkoofPage(logger))
		r.Post("/skoof/community/{id}", skoof.SkoofPage(logger))
	})
}

func New(r *chi.Mux, logger *zap.SugaredLogger, config *config.Config, authService auth.AuthService) *http.Server {

	initHandlers(r, logger, authService)

	httpSrv := &http.Server{
		Addr:              config.HTTPConfig.Addr,
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return httpSrv
}
