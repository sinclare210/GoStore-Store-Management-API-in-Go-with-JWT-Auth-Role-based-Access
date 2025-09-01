package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/models"
)

func createUser(context *gin.Context) {
	var newUser models.User

	if err := context.ShouldBindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input for creating user"})
		return
	}


	if err := newUser.CreateUser(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User Created!"})
}

func loginUser(conetxt *gin.Context){

}
