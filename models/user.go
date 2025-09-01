package models

import (
	"errors"

	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/db"
)

type User struct {
	Id       int64
	Name     string
	Email    string
	Password string
	Role     string
}

func (user *User) CreateUser() error {
	query := `
	INSERT INTO users(Name,Email,Password,Role)
	VALUES(?,?,?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("invalid statement")
	}
	defer stmt.Close()

	_, err = stmt.Exec(&user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return errors.New("invalid inputs")
	}
	return nil
}
