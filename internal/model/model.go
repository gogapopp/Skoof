package model

import "time"

type (
	// SignUpUser the user model for SignInPage handler.
	SignUpUser struct {
		UserID       int       `json:""`
		Email        string    `json:"email" validate:"required"`
		Username     string    `json:"username" validate:"required"`
		PasswordHash string    `json:"password" validate:"required"`
		CreatedAt    time.Time `json:""`
		MetaInfo     string    `json:""`
		Role         string    `json:""`
	}
	// SignInUser the user model for SignInPage handler.
	SignInUser struct {
		Email        string `json:"email"`
		Username     string `json:"username"`
		PasswordHash string `json:"password" validate:"required"`
	}
)
