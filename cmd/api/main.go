package main

import (
	"crud-api-go/config"
	handler "crud-api-go/handlers"
	"crud-api-go/repository"
	"crud-api-go/routes"
	service "crud-api-go/services"
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

	// Buat instance Handler dengan menyuntikkan 'db'
	authHandler := handler.NewAuthHandler(db)
	userHandler := handler.NewUserHandler(db)

	// Mahasiswa
	mahasiswaRepo := repository.NewMahasiswaRepository(db)
	mahasiswaService := service.NewMahasiswaService(mahasiswaRepo)
	mahasiswaHandler := handler.NewMahasiswaHandler(mahasiswaService)

	// Setup router
	r := gin.Default()

	// Route root untuk menampilkan pesan selamat datang
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Selamat datang di CRUD API Mahasiswa GoLang")
	})

	// Panggil setiap fungsi Handler
	routes.AuthRoutes(r, authHandler)
	routes.UserRoutes(r, userHandler)
	routes.MahasiswaRoutes(r, mahasiswaHandler)

	// Jalankan
	r.Run(":8081")
}
