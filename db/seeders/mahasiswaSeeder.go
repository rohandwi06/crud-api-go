package main

import (
	"log"
	"os"

	"crud-api-go/config"
	model "crud-api-go/db/models"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {

	// Load file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Gagal memuat .env: %v", err)
	}

	// Buat struct Config dan isi datanya dari .env
	cfg := config.Config{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
	}

	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		log.Println("Gagal membuka koneksi dengan database :", err)
	}

	SeedMahasiswa(db)

}

func SeedMahasiswa(db *gorm.DB) {

	// Data yang akan diisikan
	mahasiswa := []model.Mahasiswa{
		{Nama: "Rohan Dwi", Nim: "108072500017", Prodi: "S1-IF", Kelas: "IF-05-06"},
	}

	// Cek apakah data sudah ada
	var count int64
	db.Model(&model.Mahasiswa{}).Count(&count)
	if count > 0 {
		// Jika data sudah ada, maka seeder tidak dijalankan
		log.Println("Seeder tidak dijalankan karena data sudah ada.")
		return
	}

	// Memasukkan data kedalam tabel database
	result := db.Create(&mahasiswa)
	if result.Error != nil {
		log.Println("Gagal menjalankan seeder :", result.Error)
	}

	log.Printf("Berhasil seed %d data Mahasiswa.\n", result.RowsAffected)
}
