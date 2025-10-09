package seeders

import (
	model "crud-api-go/db/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	users := []model.User{
		{Username: "admin", Password: string(hashedPassword)},
	}

	return db.Create(&users).Error
}
