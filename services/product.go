package services

import (
	"errors"

	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/db"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/models"
	"gorm.io/gorm"
)

func CreateProducts(Name, Description string, Price float64, Quantity int64, User_Id uint) error {
	product := &models.Product{
		Name:        Name,
		Description: Description,
		Price:       Price,
		Quantity:    Quantity,
		UserID:      User_Id,
	}
	return db.DB.Create(&product).Error
}

func GetProducts() ([]models.Product, error) {
	var product []models.Product
	err := db.DB.Find(&product).Error
	return product, err
}

func GetProductById(Id uint) (models.Product, error) {
	var product models.Product
	  err := db.DB.First(&product, Id).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return models.Product{},errors.New("product not found")
        }
        return models.Product{},err
    }
	return product,nil
}

func DeleteProduct(Id uint) error {
	return db.DB.Delete(&models.Product{}, Id).Error
}

func UpdateProduct(Id uint) error {
	
	return db.DB.Where("id = ?", Id).Updates(&models.Product{}).Error
}
