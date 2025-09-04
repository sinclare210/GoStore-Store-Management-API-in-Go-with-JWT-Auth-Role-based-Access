package services

import (
	"errors"
	"fmt"

	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/db"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/models"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/utils"
)

func CreateUser(Name,Password,Email,Role string) error {
	hashedPassword, err := utils.HashPassword(Password)
	if err != nil {
		return errors.New("could not hash password")
	}
	user := &models.User{
		Name: Name,
		Email: Email,
		Password: hashedPassword,
		Role: Role,
	}
	return db.DB.Create(user).Error
}

func LoginUser(email,password string) (*models.User,error) {
	var user models.User

	result := db.DB.Where("Email = ?", email ).First(&user)
	if result.Error != nil {
		return  &models.User{},result.Error
	}
	fmt.Print(user)

	err := utils.CheckHashPassWord(user.Password, password)
	if err != nil {
		return &models.User{},errors.New("invalid credentials")
	}

	return &user, nil
}
