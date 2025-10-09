package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	// Fungsi di dalam ini adalah middleware yang akan dijalankan Gin
	return func(c *gin.Context) {
		// Ambil header "Authorization" dari request yang masuk
		authHeader := c.GetHeader("Authorization")

		// Cek apakah header-nya ada atau tidak
		if authHeader == "" {
			// Jika tidak ada kirim response error
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			// Hentikan proses request
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse dan validasi token yang udah diekstrak
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		// Cek hasil validasi
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort() // Hentikan proses
			return
		}

		// Lanjutkan request ke controller yang dituju
		c.Next()
	}
}
