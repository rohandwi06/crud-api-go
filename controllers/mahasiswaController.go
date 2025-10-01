package controllers

import (
	models "crud-api-go/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Membuat struct MahasiswaController sebagai blueprint
type MahasiswaController struct {
	DB *gorm.DB
}

func NewMahasiswaController(db *gorm.DB) *MahasiswaController {
	return &MahasiswaController{DB: db}
}

// == CREATE ==
func (ctrl *MahasiswaController) CreateMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctrl.DB.Create(&mahasiswa)
	c.JSON(http.StatusOK, mahasiswa)
}

// == READ ALL ==
func (ctrl *MahasiswaController) GetMahasiswas(c *gin.Context) {
	var mahasiswas []models.Mahasiswa
	ctrl.DB.Find(&mahasiswas)
	c.JSON(http.StatusOK, mahasiswas)
}

// == READ BY ID ==
func (ctrl *MahasiswaController) GetMahasiswa(c *gin.Context) {
	id := c.Param("id")
	var mahasiswa models.Mahasiswa

	if err := ctrl.DB.First(&mahasiswa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, mahasiswa)
}

// == UPDATE ==
// FUNGSI BARU UNTUK UPDATE
func (ctrl *MahasiswaController) UpdateMahasiswa(c *gin.Context) {
	id := c.Param("id")
	var mahasiswa models.Mahasiswa

	// 1. Cari dulu data mahasiswa berdasarkan ID, kalau tidak ada, kembalikan error
	if err := ctrl.DB.First(&mahasiswa, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
		return
	}

	// 2. Bind data JSON dari request body ke variabel 'input'
	var input models.Mahasiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. Update data mahasiswa yang sudah ditemukan dengan data dari 'input'
	ctrl.DB.Model(&mahasiswa).Updates(input)

	// 4. Kembalikan data mahasiswa yang sudah di-update
	c.JSON(http.StatusOK, mahasiswa)
}

// == DELETE ==
// FUNGSI BARU UNTUK DELETE
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
