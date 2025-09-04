package db

import (
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB,err = gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	err = DB.AutoMigrate(&models.User{},&models.Product{},&models.Order{})

	if err != nil {
		panic("Could not create tables")
	}

}

