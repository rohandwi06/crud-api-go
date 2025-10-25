package handler

import (
	model "crud-api-go/db/models" // Import package model untuk akses struktur User dari database
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin" // Framework web Gin
	"gorm.io/gorm"             // ORM Gorm untuk akses database

	"github.com/golang-jwt/jwt/v5" // Library untuk buat & verifikasi JWT token
	"golang.org/x/crypto/bcrypt"   // Library untuk hashing & verifikasi password
)

// Ambil secret key JWT dari environment variable
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// Struct AuthHandler menyimpan koneksi database
type AuthHandler struct {
	DB *gorm.DB
}

// Fungsi constructor untuk membuat instance AuthHandler dengan koneksi DB
func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{DB: db}
}

// Fungsi handler untuk login
func (ctrl *AuthHandler) Login(c *gin.Context) {
	// Struct sementara untuk menerima inputan dari JSON request
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind data JSON request ke struct input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Buat variabel untuk menampung data user dari database
	var user model.User

	// Cari user berdasarkan username di database
	if err := ctrl.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Verifikasi password yang dimasukkan dengan password hash di database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	// Buat JWT token berisi klaim user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,                              // ID user
		"username": user.Username,                        // Username user
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expired dalam 1 jam
	})

	// Tanda tangani token dengan secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	// Kirim respon sukses ke client berisi token
	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"token":   tokenString,
	})
}
