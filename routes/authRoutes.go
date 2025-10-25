package routes

import (
	handler "crud-api-go/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, ctrl *handler.AuthHandler) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", ctrl.Login)
	}
}
