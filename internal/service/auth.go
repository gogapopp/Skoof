package service

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/gogapopp/Skoof/internal/libs/jwt"
	"github.com/gogapopp/Skoof/internal/model"
)

func (a *authService) SignUp(ctx context.Context, user model.SignUpUser) error {
	const op = "service.auth.SignUp"
	err := a.validator.Struct(user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	err = validateRole(user.Role)
	if err != nil {
		return fmt.Errorf("%s: %w", op, ErrUndefinedRole)
	}
	user.PasswordHash = a.generatePasswordHash(user.PasswordHash)
	err = a.authStore.SignUp(ctx, user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *authService) SignIn(ctx context.Context, user model.SignInUser) (string, error) {
	const op = "service.auth.SignIn"
	err := a.validator.Struct(user)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	user.PasswordHash = a.generatePasswordHash(user.PasswordHash)
	userID, userRole, err := a.authStore.SignIn(ctx, user)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	token, err := jwt.GenerateJWTToken(userID, userRole)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

func (a *authService) generatePasswordHash(password string) string {
	hash := sha256.New()
	_, _ = hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(a.passSecret)))
}

func validateRole(role string) error {
	if role != "admin" && role != "user" {
		return ErrUndefinedRole
	}
	return nil
}
