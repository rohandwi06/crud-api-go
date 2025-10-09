package seeders

import (
	model "crud-api-go/db/models"

	"gorm.io/gorm"
)

func SeedMahasiswa(db *gorm.DB) error {

	// Data yang akan diisikan
	mahasiswa := []model.Mahasiswa{
		{Nama: "Rohan Dwi", Nim: "108072500017", Prodi: "S1-IF", Kelas: "IF-05-06"},
		{Nama: "Jane Doe", Nim: "123456789101", Prodi: "S1-IF", Kelas: "IF-05-06"},
	}

	return db.Create(&mahasiswa).Error
}
