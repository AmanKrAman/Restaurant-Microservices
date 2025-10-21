package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key")

type CustomClaims struct {
	UserID  uint   `json:"user_id"`
	Name    string `json:"name"`
	OrgName string `json:"org_name"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID uint, name string) (string, error) {
	claims := CustomClaims{
		UserID:  userID,
		Name:    name,
		OrgName: "public",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "user-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err // returns proper parsing/validation error
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
