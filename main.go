package main

import (
	"crud-api-go/config"
	"crud-api-go/controllers"
	"crud-api-go/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Muat .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal load .env :", err)
	}

	// Muat konfigurasi
	cfg := config.Config{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
	}

	// Buat koneksi database
	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Fatal("Tidak bisa konek ke database: ", err)
	}

	// Buat instance controller dengan menyuntikkan 'db'
	mahasiswaController := controllers.NewMahasiswaController(db)

	// Setup router
	r := gin.Default()

	// Panggil setiap fungsi controller
	routes.SetupRoutes(r, mahasiswaController)

	// Jalankan
	r.Run()
}
