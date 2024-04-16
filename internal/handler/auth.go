package handler

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/gogapopp/Skoof/components/auth_pages"
	"github.com/gogapopp/Skoof/internal/model"
	"github.com/gogapopp/Skoof/internal/service"
	"github.com/gogapopp/Skoof/internal/storage"
	"go.uber.org/zap"
)

var valErr validator.ValidationErrors

type authService interface {
	SignUp(ctx context.Context, user model.User) error
	SignIn(ctx context.Context, user model.SignInUser) (string, error)
}

func SignUpPage(logger *zap.SugaredLogger, a authService) http.HandlerFunc {
	const op = "handler.auth.SignUpPage"
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

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
				http.Error(w, "passwords doesn't equals", http.StatusBadRequest)
				return
			}

			err := a.SignUp(ctx, model.User{
				Username:     userName,
				Email:        userEmail,
				PasswordHash: userPassword,
				CreatedAt:    time.Now(),
				Role:         "user",
			})
			if err != nil {
				logger.Errorf("%s: %w", op, err)
				// TODO: add error component
				if errors.Is(err, storage.ErrUserExists) {
					http.Error(w, "user already exists", http.StatusBadRequest)
					return
				}
				if errors.As(err, &valErr) {
					http.Error(w, "email, password and username is required field", http.StatusBadRequest)
					return
				}
				if errors.Is(err, service.ErrUndefinedRole) {
					http.Error(w, "undefined role (available roles: admin, user)", http.StatusBadRequest)
					return
				}
				http.Error(w, "error create user", http.StatusBadRequest)
				return
			}

			http.Redirect(w, r, "/signin", http.StatusSeeOther)
			return

		case http.MethodGet:
			if err := render(r.Context(), w, auth_pages.SignUpBase(auth_pages.SignUp())); err != nil {
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

func SignInPage(logger *zap.SugaredLogger, a authService) http.HandlerFunc {
	const op = "handler.auth.SignInPage"
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		switch r.Method {
		case http.MethodPost:
			emailOrUsername := r.FormValue("email_or_username")
			userPassword := r.FormValue("password")
			// add validation
			// add session
			// update
			token, err := a.SignIn(ctx, model.SignInUser{
				Email:        emailOrUsername,
				Username:     emailOrUsername,
				PasswordHash: userPassword,
			})
			if err != nil {
				logger.Errorf("%s: %w", op, err)
				// TODO: add error component
				if errors.Is(err, storage.ErrUserNotExist) {
					http.Error(w, "invalid email/username or password", http.StatusBadRequest)
					return
				}
				if errors.As(err, &valErr) {
					http.Error(w, "email or username and password is required field", http.StatusBadRequest)
					return
				}
				http.Error(w, "something went wrong", http.StatusInternalServerError)
				return
			}

			cookie := &http.Cookie{
				Name:     "ssid",
				Value:    token,
				Path:     "/",
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
				Secure:   true}
			http.SetCookie(w, cookie)

			http.Redirect(w, r, "/skoof", http.StatusSeeOther)
			return

		case http.MethodGet:
			if err := render(ctx, w, auth_pages.SignInBase(auth_pages.SignIn())); err != nil {
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
