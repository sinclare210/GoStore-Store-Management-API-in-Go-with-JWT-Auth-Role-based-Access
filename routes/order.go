package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/models"
)

func createOrder(context *gin.Context) {
	Id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id"})
		return
	}

	var order models.Order

	order.Product_Id = Id

	user_Id, exist := context.Get("Id")
	if !exist {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "login"})
		return
	}
	_, err = models.GetProductById(Id) 
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	order.User_Id = user_Id.(int64)

	err = order.CreateOrder()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "order Created!"})

}
