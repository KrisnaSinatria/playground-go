package main

import (
	"go-first/database"
	"go-first/models"
	"go-first/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	err := database.DB.AutoMigrate(
		&models.CategoryPackage{},
		&models.Package{},
		&models.User{},
		&models.Admin{},
		&models.Staff{},
		&models.Customer{},
	)
	if err != nil {
		panic("Gagal melakukan AutoMigrate: " + err.Error())
	}

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}