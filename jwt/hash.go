package jwt

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", errors.New("failed to hash passsord")
	}
	return string(HashedPassword), nil
}

func UnHashPassWord(hashedPassword,passsord string)(error){
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(passsord))
	if err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}
