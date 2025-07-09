package main

import (
	"candra/backend-api/config"
	"candra/backend-api/database"
	"candra/backend-api/routes"
)

func main() {
	config.LoadEnv()                               // Muat konfigurasi dari .env
	database.InitDatabase()                        // Inisialisasi database
	r := routes.SetupRouter()                      // Setup router
	r.Run(":" + config.GetEnv("APP_PORT", "3000")) // Jalankan server
}
