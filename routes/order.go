package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/db"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/models"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/services"
)

func createOrder(context *gin.Context) {
	Id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id"})
		return
	}

	var order models.Order

	order.ProductID = uint(Id)

	user_Id, exist := context.Get("Id")
	if !exist {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "login"})
		return
	}
	product, err := services.GetProductById(uint(Id))
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	fmt.Println(product)

	order.ProductName = product.Name
	order.ProductPrice = product.Price
	order.UserID = uint(user_Id.(int64))

	err = services.CreateOrder(order.UserID, order.ProductID, order.ProductName, order.ProductPrice)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "order Created!"})

}

func getOrderByUser(context *gin.Context) {
	user_Id, exist := context.Get("Id")
	if !exist {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "login"})
		return
	}

	orders, err := services.GetOrdersForUser(uint(user_Id.(int64)))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": orders})
}

func deleteOrder(context *gin.Context) {
   
    Id, err := strconv.ParseInt(context.Param("id"), 10, 64)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id"})
        return
    }

   
    userId, exist := context.Get("Id")
    if !exist {
        context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
        return
    }


    var order models.Order
    if err := db.DB.First(&order, uint(Id)).Error; err != nil {
        context.JSON(http.StatusNotFound, gin.H{"message": "Order not found"})
        return
    }


    role, _ := context.Get("Role")
    if role != "admin" && order.UserID != uint(userId.(int64)) {
        context.JSON(http.StatusForbidden, gin.H{"message": "Not allowed to delete this order"})
        return
    }


    result := db.DB.Delete(&order)
    if result.Error != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
        return
    }
    if result.RowsAffected == 0 {
        context.JSON(http.StatusNotFound, gin.H{"message": "Order not found"})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

