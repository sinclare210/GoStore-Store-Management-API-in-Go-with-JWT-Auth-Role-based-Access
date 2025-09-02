package routes

import (
	"net/http"
	"strconv"

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

func getProducts(context *gin.Context) {

	products, err := models.GetProducts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": products})
}

func getProduct(context *gin.Context) {
	Id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id"})
		return
	}

	product, err := models.GetProductById(Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": product})

}

func deleteProduct(context *gin.Context) {
	Id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id"})
		return
	}

	product, err := models.GetProductById(Id) 
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
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

	productId, exist := context.Get("Id")
	if !exist {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Role"})
		return
	}

	if productId != product.User_Id {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not yours to delete"})
		return
	}

	err = product.DeleteProduct()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Product Deleted!"})

}

func updateProduct(context *gin.Context) {
	var updateProduct models.Product

	err := context.ShouldBindJSON(&updateProduct)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input for creating user"})
		return
	}

	Id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Id"})
		return
	}

	product, err := models.GetProductById(Id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	updateProduct.Id = product.Id
	updateProduct.User_Id = product.User_Id

	Role, exist := context.Get("Role")
	if !exist {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Role"})
		return
	}

	if Role != "admin" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unautorized"})
		return
	}

	userId, exist := context.Get("Id")
	if !exist {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Role"})
		return
	}

	if userId != updateProduct.User_Id {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not yours to update"})
		return
	}

	err = updateProduct.UpdateProduct()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Product Updated!"})

}
