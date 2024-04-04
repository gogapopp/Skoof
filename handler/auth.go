package handler

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gogapopp/Skoof/components/auth_pages"
	"github.com/gogapopp/Skoof/model"
	"github.com/gogapopp/Skoof/storage"
)

func SignInPage(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			emailOrUsername := r.FormValue("email_or_username")
			userPassword := r.FormValue("password")
			log.Print(emailOrUsername, userPassword)
			// add validation
			// add session
			// update
			user, err := s.GetUser(r.Context(), emailOrUsername, generatePasswordHash(userPassword))
			if err != nil {
				// add error component
				if errors.Is(storage.ErrUserNotExist, err) {
					http.Error(w, "invalid email/username or password", http.StatusBadRequest)
					return
				}
				http.Error(w, "something went wrong", http.StatusInternalServerError)
				return
			}

			_ = user
			// TODO:
			jwtToken := ""
			w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
			http.Redirect(w, r, "/skoof", http.StatusSeeOther)
			return

		case http.MethodGet:
			if err := render(r.Context(), w, auth_pages.SignInBase(auth_pages.SignIn())); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func SignUpPage(s Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// add validation
			// add session
			// update
			userName := r.FormValue("username")
			userEmail := r.FormValue("email")
			userPassword := r.FormValue("password")
			userPasswordConfirm := r.FormValue("password_confirm")

			if userPassword != userPasswordConfirm {
				// TODO: add error component
				http.Error(w, "password doesn't equals", http.StatusBadRequest)
				return
			}

			user := model.User{
				Username:  userName,
				Email:     userEmail,
				Password:  generatePasswordHash(userPassword),
				CreatedAt: time.Now(),
			}

			err := s.CreateUser(r.Context(), user)
			if err != nil {
				// TODO: add error component
				http.Error(w, "error create user", http.StatusBadRequest)
				return
			}

			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return

		case http.MethodGet:
			if err := render(r.Context(), w, auth_pages.SignUpBase(auth_pages.SignUp())); err != nil {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

func generatePasswordHash(password string) string {
	// TODO:
	const secretKey = "secret_key"
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(secretKey)))
}
