package routes

import (
	controller "crud-api-go/controllers"
	"crud-api-go/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, ctrl *controller.UserController) {
	api := r.Group("/api", middleware.AuthMiddleware())
	{
		api.GET("/users", ctrl.GetAllUsers)
		api.GET("/users/:id", ctrl.GetUserById)
		api.POST("/users", ctrl.CreateUser)
		api.PUT("/users/:id", ctrl.UpdateUser)
		api.DELETE("/users/:id", ctrl.DeleteUser)
	}
}
