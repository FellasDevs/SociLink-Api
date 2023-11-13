package routes

import (
	postcontroller "SociLinkApi/controllers/post"
	"SociLinkApi/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TimelineRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("", middlewares.AuthenticateUser, func(context *gin.Context) {
		context.JSON(http.StatusNotImplemented, gin.H{
			"success": false,
			"message": "Rota não implementada",
		})
	})

	router.GET("/:id", func(context *gin.Context) {
		postcontroller.GetUserTimeline(context, db)
	})
}
