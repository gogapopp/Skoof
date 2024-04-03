package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gogapopp/Skoof/components"
)

var (
	errUsrNotExst     = errors.New("user not exists")
	errUsrAlreadyExst = errors.New("user already exists")
)

func IndexPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := render(r.Context(), w, components.IndexPage("Skoof")); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		// w.WriteHeader(http.StatusOK)
		// fmt.Fprint(w, "WORLD WIDE SKOOF")
	}
}

func SignInPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userLogin := r.FormValue("email")
			userPassword := r.FormValue("password")
			log.Print(userLogin, userPassword)
			// add validation
			// add session
			// update
			if userLogin != "" || userPassword != "" {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

		case http.MethodGet:
			if err := render(r.Context(), w, components.LoginPageBase(components.LoginPage())); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func SignUpPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// add validation
			// add session
			// update
			userLogin := r.FormValue("login")
			userPassword := r.FormValue("password")
			userPasswordConfirm := r.FormValue("password")

			log.Print(userLogin, userPassword, userPasswordConfirm)

			err := stubGetUser()
			if err == errUsrAlreadyExst {
				http.Redirect(w, r, "/signin", http.StatusTemporaryRedirect)
				return
			}

		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, "SIGNUP")

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func stubGetUser() error {
	return errUsrNotExst
}
