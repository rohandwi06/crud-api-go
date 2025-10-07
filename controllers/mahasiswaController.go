package controllers

import (
	models "crud-api-go/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Membuat struct MahasiswaController berisi blueprint koneksi database
type MahasiswaController struct {
	DB *gorm.DB
}

// Assign koneksi database pada struct MahasiswaController
func NewMahasiswaController(db *gorm.DB) *MahasiswaController {
	return &MahasiswaController{DB: db}
}

// CREATE MAHASISWA
func (ctrl *MahasiswaController) CreateMahasiswa(c *gin.Context) {

	// Deklarasi variabel sebagai blueprint data mahasiswa
	var mahasiswa models.Mahasiswa

	// Mengambil input dari user dan menyimpannya di variabel mahasisw
	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah nim sudah terdaftar
	var nimTerdaftar models.Mahasiswa
	result := ctrl.DB.Where("nim = ?", mahasiswa.Nim).First(&nimTerdaftar)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "NIM sudah terdaftar!",
		})
		return
	}

	// Buat data mahasiswa baru
	ctrl.DB.Create(&mahasiswa)

	// Return data mahasiswa yang sudah dibuat
	c.JSON(http.StatusCreated, gin.H{
		"data":    mahasiswa,
		"message": "Data berhasil dibuat!",
	})
}

// READ ALL MAHASISWA
func (ctrl *MahasiswaController) GetAllMahasiswa(c *gin.Context) {

	// Deklarasi variabel dengan blueprint untuk menyimpan banyak data mahasiswa
	var mahasiswas []models.Mahasiswa

	// Mengambil semua data mahasiswa yang ada
	ctrl.DB.Find(&mahasiswas)

	// Return data mahasiswa yang diambil
	c.JSON(http.StatusOK, gin.H{
		"data":    mahasiswas,
		"message": "Data berhasil didapatkan!",
	})
}

// READ MAHASISWA BY ID
func (ctrl *MahasiswaController) GetMahasiswaById(c *gin.Context) {

	// Mengambil ID dari parameter
	id := c.Param("id")

	// Deklarasi variabel dengan blueprint untuk menyimpan data mahasiswa
	var mahasiswa models.Mahasiswa

	// Cek data mahasiswa berdasarkan ID
	if err := ctrl.DB.First(&mahasiswa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// UPDATE MAHASISWA
func (ctrl *MahasiswaController) UpdateMahasiswa(c *gin.Context) {

	// Ambil id dari parameter
	id := c.Param("id")

	// Deklarasi variabel sebagai blueprint untuk menyimpan data mahasiswa
	var mahasiswa models.Mahasiswa

	// Cek data mahasiswa berdasarkan ID
	if err := ctrl.DB.First(&mahasiswa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
		return
	}

	// Ambil data input dari user
	var input models.Mahasiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update data mahasiswa yang sudah ditemukan dengan data dari input
	ctrl.DB.Model(&mahasiswa).Updates(input)

	// Return data mahasiswa yang sudah di update
	c.JSON(http.StatusOK, mahasiswa)
}

// DELETE MAHASISWA
func (ctrl *MahasiswaController) DeleteMahasiswa(c *gin.Context) {
	id := c.Param("id")
	var mahasiswa models.Mahasiswa

	// 1. Cari dulu datanya, kalau ngga ada, kasih error
	if err := ctrl.DB.First(&mahasiswa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	// 2. Hapus data mahasiswa dari database
	ctrl.DB.Delete(&mahasiswa)

	// 3. Kembalikan pesan sukses
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
