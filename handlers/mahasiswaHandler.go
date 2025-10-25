package handler

import (
	models "crud-api-go/db/models"
	service "crud-api-go/services" // Import service lu
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Struct-nya sekarang megang Service, BUKAN gorm.DB
type MahasiswaHandler struct {
	Service service.MahasiswaService
}

// Constructor-nya juga nerima Service
func NewMahasiswaHandler(service service.MahasiswaService) *MahasiswaHandler {
	return &MahasiswaHandler{Service: service}
}

// CREATE MAHASISWA
func (ctrl *MahasiswaHandler) CreateMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa

	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Panggil Service
	createdMahasiswa, err := ctrl.Service.CreateMahasiswa(mahasiswa)
	if err != nil {
		// Handler nerjemahin error dari service ke HTTP status
		if err.Error() == "NIM_EXISTS" {
			c.JSON(http.StatusConflict, gin.H{"error": "NIM sudah terdaftar!"})
			return
		}

		// Error internal server
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    createdMahasiswa,
		"message": "Data berhasil dibuat!",
	})
}

// READ ALL MAHASISWA
func (ctrl *MahasiswaHandler) GetAllMahasiswa(c *gin.Context) {
	mahasiswas, err := ctrl.Service.GetAllMahasiswa()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    mahasiswas,
		"message": "Data berhasil didapatkan!",
	})
}

// READ MAHASISWA BY ID
func (ctrl *MahasiswaHandler) GetMahasiswaById(c *gin.Context) {
	id := c.Param("id")

	mahasiswa, err := ctrl.Service.GetMahasiswaById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ini yang kurang di kode asli lu:
	c.JSON(http.StatusOK, gin.H{
		"data":    mahasiswa,
		"message": "Data berhasil didapatkan!",
	})
}

// UPDATE MAHASISWA
func (ctrl *MahasiswaHandler) UpdateMahasiswa(c *gin.Context) {
	id := c.Param("id")

	var input models.Mahasiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedMahasiswa, err := ctrl.Service.UpdateMahasiswa(id, input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Mahasiswa tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return data yg udah di-update
	c.JSON(http.StatusOK, gin.H{
		"data":    updatedMahasiswa,
		"message": "Data berhasil diupdate!",
	})
}

// DELETE MAHASISWA
func (ctrl *MahasiswaHandler) DeleteMahasiswa(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.Service.DeleteMahasiswa(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
