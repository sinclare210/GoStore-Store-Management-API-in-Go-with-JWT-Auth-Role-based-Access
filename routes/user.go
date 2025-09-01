package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/models"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/utils"
)

func createUser(context *gin.Context) {
	var newUser models.User

	if err := context.ShouldBindJSON(&newUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input for creating user"})
		return
	}

	err := newUser.CreateUser()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User Created!"})
}

func loginUser(context *gin.Context) {
	var loginUser models.User

	err := context.ShouldBindJSON(&loginUser)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input for loggin user"})
		return
	}

	err = loginUser.LoginUser(loginUser.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(loginUser.Email, loginUser.Id, loginUser.Role)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Login Successful!", "token": token, "Role": loginUser.Role})
}
