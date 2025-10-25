package repository

import (
	models "crud-api-go/db/models"

	"gorm.io/gorm"
)

type MahasiswaRepository interface {
	Save(mahasiswa *models.Mahasiswa) error
	FindAll() ([]models.Mahasiswa, error)
	FindById(id string) (models.Mahasiswa, error)
	FindByNIM(nim string) (models.Mahasiswa, error)
	Update(mahasiswa *models.Mahasiswa) error
	Delete(mahasiswa *models.Mahasiswa) error
}

type mahasiswaRepository struct {
	DB *gorm.DB
}

func NewMahasiswaRepository(db *gorm.DB) MahasiswaRepository {
	return &mahasiswaRepository{DB: db}
}

func (r *mahasiswaRepository) Save(mahasiswa *models.Mahasiswa) error {
	return r.DB.Create(mahasiswa).Error
}

func (r *mahasiswaRepository) FindAll() ([]models.Mahasiswa, error) {
	var mahasiswas []models.Mahasiswa
	err := r.DB.Find(&mahasiswas).Error
	return mahasiswas, err
}

func (r *mahasiswaRepository) FindById(id string) (models.Mahasiswa, error) {
	var mahasiswa models.Mahasiswa
	err := r.DB.First(&mahasiswa, id).Error
	return mahasiswa, err
}

func (r *mahasiswaRepository) FindByNIM(nim string) (models.Mahasiswa, error) {
	var mahasiswa models.Mahasiswa
	err := r.DB.Where("nim = ?", nim).First(&mahasiswa).Error
	return mahasiswa, err
}

func (r *mahasiswaRepository) Update(mahasiswa *models.Mahasiswa) error {
	// GORM's Updates() butuh model yg ada ID-nya
	// Kita asumsikan 'mahasiswa' yg dipassing udah berisi ID yg bener
	return r.DB.Save(mahasiswa).Error
	// Atau pakai .Updates() jika hanya field tertentu
	// return r.DB.Model(mahasiswa).Updates(mahasiswa).Error
}

func (r *mahasiswaRepository) Delete(mahasiswa *models.Mahasiswa) error {
	return r.DB.Delete(mahasiswa).Error
}
