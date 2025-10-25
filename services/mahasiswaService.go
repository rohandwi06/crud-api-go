package service

import (
	models "crud-api-go/db/models"
	"crud-api-go/repository" // Import repository lu
	"errors"                 // Kita butuh ini untuk custom error

	"gorm.io/gorm"
)

// Interface
type MahasiswaService interface {
	CreateMahasiswa(input models.Mahasiswa) (models.Mahasiswa, error)
	GetAllMahasiswa() ([]models.Mahasiswa, error)
	GetMahasiswaById(id string) (models.Mahasiswa, error)
	UpdateMahasiswa(id string, input models.Mahasiswa) (models.Mahasiswa, error)
	DeleteMahasiswa(id string) error
}

// Struct implementasi
type mahasiswaService struct {
	repo repository.MahasiswaRepository // Dependency-nya ke interface Repo
}

// Constructor
func NewMahasiswaService(repo repository.MahasiswaRepository) MahasiswaService {
	return &mahasiswaService{repo: repo}
}

// Implementasi method

func (s *mahasiswaService) CreateMahasiswa(input models.Mahasiswa) (models.Mahasiswa, error) {
	// === INI BUSINESS LOGIC ===
	// Cek apakah nim sudah terdaftar
	_, err := s.repo.FindByNIM(input.Nim)

	// Kalo `err == nil` berarti data DITEMUKAN, alias NIM udah ada
	if err == nil {
		return models.Mahasiswa{}, errors.New("NIM_EXISTS")
	}

	// Kalo error-nya BUKAN "data tidak ditemukan", berarti ada error database lain
	if err != gorm.ErrRecordNotFound {
		return models.Mahasiswa{}, err // Error DB (e.g., connection lost)
	}
	// === SELESAI BUSINESS LOGIC ===

	// Kalo lolos, baru panggil repository untuk save
	err = s.repo.Save(&input)
	if err != nil {
		return models.Mahasiswa{}, err
	}

	return input, nil
}

func (s *mahasiswaService) GetAllMahasiswa() ([]models.Mahasiswa, error) {
	// Di sini gaada business logic, jadi tinggal lempar aja
	return s.repo.FindAll()
}

func (s *mahasiswaService) GetMahasiswaById(id string) (models.Mahasiswa, error) {
	// Tinggal lempar. Nanti repo yg return error gorm.ErrRecordNotFound
	return s.repo.FindById(id)
}

func (s *mahasiswaService) UpdateMahasiswa(id string, input models.Mahasiswa) (models.Mahasiswa, error) {
	// Cek dulu datanya ada apa ngga
	mahasiswa, err := s.repo.FindById(id)
	if err != nil {
		return models.Mahasiswa{}, err // Bakal return gorm.ErrRecordNotFound kalo gaada
	}

	// Update datanya
	mahasiswa.Nama = input.Nama
	mahasiswa.Nim = input.Nim
	mahasiswa.Prodi = input.Prodi
	mahasiswa.Kelas = input.Kelas

	// Simpen perubahan
	err = s.repo.Update(&mahasiswa)
	if err != nil {
		return models.Mahasiswa{}, err
	}

	return mahasiswa, nil
}

func (s *mahasiswaService) DeleteMahasiswa(id string) error {
	// Cek dulu datanya
	mahasiswa, err := s.repo.FindById(id)
	if err != nil {
		return err // Data not found
	}

	// Kalo ada, hapus
	return s.repo.Delete(&mahasiswa)
}
