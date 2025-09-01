package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/models"
)

func createProduct(context *gin.Context) {
	var product models.Product

	err := context.ShouldBindJSON(&product)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input for creating user"})
		return
	}

	Role, exist := context.Get("Role")
	if !exist {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Role"})
		return
	}

	if Role != "admin" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unautorized"})
		return
	}

	Id, _ := context.Get("Id")

	product.User_Id = Id.(int64)

	err = product.CreateProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Product Created!"})

}
