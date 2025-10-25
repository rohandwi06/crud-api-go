package routes

import (
	handler "crud-api-go/handlers"
	"crud-api-go/middleware"

	"github.com/gin-gonic/gin"
)

func MahasiswaRoutes(router *gin.Engine, ctrl *handler.MahasiswaHandler) {

	api := router.Group("/api", middleware.AuthMiddleware())
	{
		api.POST("/mahasiswa", ctrl.CreateMahasiswa)
		api.GET("/mahasiswa", ctrl.GetAllMahasiswa)
		api.GET("/mahasiswa/:id", ctrl.GetMahasiswaById)
		api.PUT("/mahasiswa/:id", ctrl.UpdateMahasiswa)
		api.DELETE("/mahasiswa/:id", ctrl.DeleteMahasiswa)
	}
}
