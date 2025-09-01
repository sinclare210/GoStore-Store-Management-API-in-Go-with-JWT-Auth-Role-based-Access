package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sinclare210/GoStore-Store-Management-API-in-Go-with-JWT-Auth-Role-based-Access/routes"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":8080")

}
