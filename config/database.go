package config

import (
	"fmt"

	"gorm.io/driver/mysql" // Dependensi driver untuk koneki dengan database
	"gorm.io/gorm"         // Dependensi untuk implementasi ORM
)

func ConnectDatabase(cfg Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Gagal konek database:", err)
		return nil, err
	}

	fmt.Println("Database connected!")
	return database, nil
}
