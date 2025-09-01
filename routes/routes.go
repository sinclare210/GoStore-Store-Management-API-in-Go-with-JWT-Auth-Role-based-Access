package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", createUser)
	server.POST("/login", loginUser)
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/product", createProduct)
	server.GET("/product", getProducts)
	server.GET("/product/:id", getProduct)
}
