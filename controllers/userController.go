package controller

import (
	models "crud-api-go/db/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Membuat struct UserController berisi blueprint koneksi database
type UserController struct {
	DB *gorm.DB
}

// Assign koneksi database pada struct UserController
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// CREATE USER
func (ctrl *UserController) CreateUser(c *gin.Context) {

	// Deklarasi variabel sebagai blueprint data user
	var user models.User

	// Mengambil input dari user dan menyimpannya di variabel user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah username sudah terdaftar
	var userTerdaftar models.User
	result := ctrl.DB.Where("username = ?", user.Username).First(&userTerdaftar)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Username sudah terdaftar!",
		})
		return
	}

	// Hash password sebelum disimpan ke database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}
	user.Password = string(hashedPassword)

	// Buat data user baru
	ctrl.DB.Create(&user)

	// Return data user yang sudah dibuat
	c.JSON(http.StatusCreated, gin.H{
		"data":    user,
		"message": "User berhasil dibuat!",
	})
}

// READ ALL USER
func (ctrl *UserController) GetAllUsers(c *gin.Context) {

	// Deklarasi variabel dengan blueprint untuk menyimpan banyak data user
	var users []models.User

	// Mengambil semua data user yang ada
	ctrl.DB.Find(&users)

	// Return data user yang diambil
	c.JSON(http.StatusOK, gin.H{
		"data":    users,
		"message": "Data user berhasil didapatkan!",
	})
}

// READ USER BY ID
func (ctrl *UserController) GetUserById(c *gin.Context) {

	// Mengambil ID dari parameter
	id := c.Param("id")

	// Deklarasi variabel dengan blueprint untuk menyimpan data user
	var user models.User

	// Cek data user berdasarkan ID
	if err := ctrl.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Return data user yang ditemukan
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "User berhasil ditemukan!",
	})
}

// UPDATE USER
func (ctrl *UserController) UpdateUser(c *gin.Context) {

	// Ambil id dari parameter
	id := c.Param("id")

	// Deklarasi variabel sebagai blueprint untuk menyimpan data user
	var user models.User

	// Cek data user berdasarkan ID
	if err := ctrl.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Ambil data input dari user
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Jika password ikut diupdate, hash ulang password-nya
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengenkripsi password"})
			return
		}
		input.Password = string(hashedPassword)
	}

	// Update data user dengan input baru
	ctrl.DB.Model(&user).Updates(input)

	// Return data user yang sudah diupdate
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "User berhasil diupdate!",
	})
}

// DELETE USER
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	// 1. Cari dulu datanya, kalau tidak ada, kasih error
	if err := ctrl.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// 2. Hapus data user dari database
	ctrl.DB.Delete(&user)

	// 3. Kembalikan pesan sukses
	c.JSON(http.StatusOK, gin.H{"message": "User berhasil dihapus"})
}
