package model

import "time"

type (
	// User the user model for SignUpPage and SignInPage handlers.
	User struct {
		UserID       int       `json:""`
		Email        string    `json:"email" validate:"required"`
		Username     string    `json:"username" validate:"required"`
		PasswordHash string    `json:"password" validate:"required"`
		CreatedAt    time.Time `json:""`
		MetaInfo     string    `json:""`
		Role         string    `json:""`
	}
	SignInUser struct {
		Email        string `json:"email"`
		Username     string `json:"username"`
		PasswordHash string `json:"password" validate:"required"`
	}
)
