package routes

import (
	controller "crud-api-go/controllers"

	"github.com/gin-gonic/gin"
)

func MahasiswaRoutes(router *gin.Engine, mahasiswaController *controller.MahasiswaController) {
	router.POST("/mahasiswa", mahasiswaController.CreateMahasiswa)
	router.GET("/mahasiswa", mahasiswaController.GetAllMahasiswa)
	router.GET("/mahasiswa/:id", mahasiswaController.GetMahasiswaById)
	router.PUT("/mahasiswa/:id", mahasiswaController.UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id", mahasiswaController.DeleteMahasiswa)
}
