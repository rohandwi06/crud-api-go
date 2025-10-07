package routes

import (
	"crud-api-go/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes akan mendaftarkan semua route dari controller ke router engine
func SetupRoutes(router *gin.Engine, mahasiswaController *controllers.MahasiswaController) {
	// Definisikan semua route di sini
	router.POST("/mahasiswa", mahasiswaController.CreateMahasiswa)
	router.GET("/mahasiswa", mahasiswaController.GetAllMahasiswa)
	router.GET("/mahasiswa/:id", mahasiswaController.GetMahasiswaById)
	router.PUT("/mahasiswa/:id", mahasiswaController.UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id", mahasiswaController.DeleteMahasiswa)
}
