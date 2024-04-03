package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/gogapopp/Skoof/handler"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handler.IndexPage())

	r.Get("/signin", handler.SignInPage())
	r.Post("/signin", handler.SignInPage())

	r.Get("/signup", handler.SignUpPage())
	r.Post("/signup", handler.SignUpPage())

	srv := http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Fatal(srv.ListenAndServe())
}
