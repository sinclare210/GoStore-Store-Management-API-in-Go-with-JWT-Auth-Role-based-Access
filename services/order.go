package services

import (
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/db"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/models"
)

func CreateOrder(User_Id, Product_Id uint, Name string, Price float64) error {
	order := &models.Order{
		UserID:       User_Id,
		ProductID:    Product_Id,
		ProductName:  Name,
		ProductPrice: Price,
	}
	return db.DB.Create(order).Error
}

func GetOrdersForUser(id uint) ([]models.Order, error) {
	var orders []models.Order

	err := db.DB.Where("user_id = ?", id).Find(&orders).Error
	return orders, err

}
