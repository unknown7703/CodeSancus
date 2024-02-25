package routes

import (
	"project/CodeSancus-Backend/controllers"

	"github.com/gin-gonic/gin"
)

func DatasetRoutes(ctx *gin.Engine) {
	admin := ctx.Group("/admin")
	{
		admin.POST("/analyse/", controllers.Show)
	}
}
