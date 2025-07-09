package database

import (
	"fmt"
	"log"
	"candra/backend-api/config"
	"candra/backend-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	// Ambil konfigurasi database dari .env
	dbUser := config.GetEnv("DB_USER", "root")
	dbPass := config.GetEnv("DB_PASS", "")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "5432")
	dbName := config.GetEnv("DB_NAME", "")

	// Buat DSN Postgres
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbName, dbPort,
	)

	// Koneksi ke database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	fmt.Println("Berhasil terhubung ke database!")

	// Migrasi otomatis tabel User dan Role
	if err := DB.AutoMigrate(&models.User{}, &models.Role{}); err != nil {
		log.Fatal("Gagal migrasi database:", err)
	}
	fmt.Println("Migrasi database berhasil!")
}
