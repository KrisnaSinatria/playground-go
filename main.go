package main

import (
	"log"

	"go-first/app/Shared/Database"
	"go-first/app/Shared/Router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Peringatan: File .env tidak ditemukan, menggunakan environment OS bawaan")
	}

	db := database.Connect()

	r := router.SetupRouter(db)

	log.Println("Server REST API berjalan di port :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
