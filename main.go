package main

import (
	"go-first/database"
	"go-first/models"
	"go-first/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan environment OS bawaan")
	}

	database.Connect()

	err := database.DB.AutoMigrate(
		&models.Room{},
	)
	if err != nil {
		panic("Gagal melakukan AutoMigrate: " + err.Error())
	}

	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
