package main

import (
	"crud-api-go/config"
	"crud-api-go/db/seeders"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal memuat .env :", err)
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
		log.Fatal("Gagal untuk koneksi db :", err)
	}

	if err := seeders.SeedUsers(db); err != nil {
		fmt.Println("Gagal seed users :", err)
	}

	if err := seeders.SeedMahasiswa(db); err != nil {
		fmt.Println("Gagal seeding mahasiswa :", err)
	}

	fmt.Println("Semua data berhasil di seed!")
}
