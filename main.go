package main

import (
	"go-crud/database"
	"go-crud/models"
	"go-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Connect()
	database.DB.AutoMigrate(&models.User{})

	routes.UserRoutes(r)

	r.Run(":8080")
}
