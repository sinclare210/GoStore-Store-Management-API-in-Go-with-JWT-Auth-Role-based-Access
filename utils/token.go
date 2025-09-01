package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecre133t"

func GenerateToken(email string, id int64, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": email,
		"Id":    id,
		"Role":  role,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

// func ValidateToken (token string) {
