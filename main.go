package main

import (
	"crud-api-go/config"
	controller "crud-api-go/controllers"
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
		log.Fatal("Gagal untuk koneksi db: ", err)
	}

	// Buat instance controller dengan menyuntikkan 'db'
	authController := controller.NewAuthController(db)
	userController := controller.NewUserController(db)
	mahasiswaController := controller.NewMahasiswaController(db)

	// Setup router
	r := gin.Default()

	// Panggil setiap fungsi controller
	routes.AuthRoutes(r, authController)
	routes.UserRoutes(r, userController)
	routes.MahasiswaRoutes(r, mahasiswaController)

	// Jalankan
	r.Run()
}
