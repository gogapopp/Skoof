package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gogapopp/Skoof/internal/config"
	"github.com/gogapopp/Skoof/internal/handler/middlewares"
	"go.uber.org/zap"
)

const readHeaderTimeout = 10 * time.Second

func initHandlers(r *chi.Mux, logger *zap.SugaredLogger, authService authService) {
	r.Get("/", HomePage(logger))

	r.Get("/signin", SignInPage(logger, authService))
	r.Post("/signin", SignInPage(logger, authService))

	r.Get("/signup", SignUpPage(logger, authService))
	r.Post("/signup", SignUpPage(logger, authService))

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)

		r.Get("/skoof", SkoofPage(logger))
		r.Post("/skoof", SkoofPage(logger))
	})
}

func New(r *chi.Mux, logger *zap.SugaredLogger, authService authService, config *config.Config) *http.Server {
	initHandlers(r, logger, authService)

	httpSrv := &http.Server{
		Addr:              config.HTTPConfig.Addr,
		Handler:           r,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	return httpSrv
}
