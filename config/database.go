package config

import (
	"fmt"

	"gorm.io/driver/mysql" // Dependensi driver untuk koneki dengan database
	"gorm.io/gorm"         // Dependensi untuk implementasi ORM
)

// Function ConnectDatabase untuk membuka koneksi dengan database
func ConnectDatabase(cfg Config) (*gorm.DB, error) {

	// Dsn atau Data Source Name adalah data yang digunakan untuk membuka koneksi dengan database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	// Membuat variabel database
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // Jika koneksi berhasil dibuat, disimpan di database

	if err != nil { // Jika koneksi gagal dibuat, maka mengembalikan nilai error
		fmt.Println("Gagal konek database:", err)
		return nil, err
	}

	fmt.Println("Database connected!")
	return database, nil
}
