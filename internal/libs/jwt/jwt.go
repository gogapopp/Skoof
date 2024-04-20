package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	errUnknownClaimsType = errors.New("unknown claims type")
)

// TODO:
const (
	SECRET_KEY = "secret_key"
	TOKEN_EXP  = time.Hour * 3
)

type tokenClaims struct {
	Role   string
	UserID int
	jwt.RegisteredClaims
}

func GenerateJWTToken(userID int, role string) (string, error) {
	claims := tokenClaims{
		role,
		userID,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXP)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ParseJWTToken(userJWTToken string) (int, string, error) {
	token, err := jwt.ParseWithClaims(userJWTToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte([]byte(SECRET_KEY)), nil
	})
	if err != nil {
		return 0, "", err
	} else if claims, ok := token.Claims.(*tokenClaims); ok {
		return claims.UserID, claims.Role, nil
	}
	return 0, "", errUnknownClaimsType
}
