package database

import (
	"fmt"
	"log"
	"os"

	itemModels "go-first/app/Domains/Item/Models"
	repairModels "go-first/app/Domains/Repair/Models"
	stockModels "go-first/app/Domains/StockMovement/Models"
	techModels "go-first/app/Domains/Technician/Models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "root")
	dbname := getEnv("DB_NAME", "system_service_electronic")
	sslmode := getEnv("DB_SSLMODE", "disable")
	timezone := getEnv("DB_TZ", "Asia/Makassar")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbname, port, sslmode, timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal connect ke database PostgreSQL: %v", err)
	}

	log.Println("Koneksi database PostgreSQL berhasil!")

	if err := db.AutoMigrate(
		&techModels.Technician{},
		&itemModels.Item{},
		&repairModels.Repair{},
		&stockModels.StockMovement{},
	); err != nil {
		log.Fatalf("Gagal menjalankan AutoMigrate database: %v", err)
	}

	log.Println("GORM AutoMigrate berhasil dilakukan untuk semua model domain.")

	DB = db
	return db
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok && val != "" {
		return val
	}
	return fallback
}
