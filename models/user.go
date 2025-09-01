package models

import (
	"errors"

	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/db"

	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/utils"
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
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {

		return errors.New("could not hash password")
	}

	user.Password = hashedPassword

	_, err = stmt.Exec(user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		return errors.New("invalid inputs")
	}
	return nil
}

func (user *User) LoginUser(password string) error {
	query := `
	SELECT Id, Email,Password,Role FROM users WHERE Email = ?
	`

	row := db.DB.QueryRow(query, user.Email)
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return errors.New("invalid query or user not found")
	}

	err = utils.CheckHashPassWord(user.Password, password)
	if err != nil {
		return errors.New("invalid credentials")
	}

	return nil
}
