package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/gogapopp/Skoof/internal/handler"
	"github.com/gogapopp/Skoof/internal/handler/middlewares"
	"github.com/gogapopp/Skoof/internal/service"
	"github.com/gogapopp/Skoof/internal/storage/sqlite"
)

const addr = ":8080"

func main() {
	db, err := sqlite.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := service.New(db)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handler.HomePage())

	r.Get("/signin", handler.SignInPage(service))
	r.Post("/signin", handler.SignInPage(service))

	r.Get("/signup", handler.SignUpPage(service))
	r.Post("/signup", handler.SignUpPage(service))

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)

		r.Get("/skoof", handler.SkoofPage())
		r.Post("/skoof", handler.SkoofPage())
	})

	srv := http.Server{
		Addr:    addr,
		Handler: r,
	}
	log.Println("server running at", addr)
	log.Fatal(srv.ListenAndServe())
}
