package routes

import (
	controller "crud-api-go/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine, ctrl *controller.AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", ctrl.Login)
	}
}
