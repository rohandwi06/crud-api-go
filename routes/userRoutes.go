package routes

import (
	handler "crud-api-go/handlers"
	"crud-api-go/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, ctrl *handler.UserHandler) {
	api := router.Group("/api", middleware.AuthMiddleware())
	{
		api.GET("/users", ctrl.GetAllUsers)
		api.GET("/users/:id", ctrl.GetUserById)
		api.POST("/users", ctrl.CreateUser)
		api.PUT("/users/:id", ctrl.UpdateUser)
		api.DELETE("/users/:id", ctrl.DeleteUser)
	}
}
