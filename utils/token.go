package utils

import (
	"errors"
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

func VerifyToken(token string) (string, int64, error) {
	pardsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return "", errors.New("invalid signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return "", 0, errors.New("wrong token parsed")
	}

	isValid := pardsedToken.Valid
	if !isValid {
		return "", 0, errors.New("invalid token parsed")
	}

	claims, ok := pardsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, errors.New("invalid claim")
	}

	Id := int64(claims["Id"].(float64))
	Role := claims["Role"].(string)

	return Role, Id, nil

}
