package routes

import (
	postcontroller "SociLinkApi/controllers/post"
	"net/http"

	"SociLinkApi/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("", middlewares.AuthenticateUser, func(context *gin.Context) {
		postcontroller.CreatePost(context, db)
	})

	router.GET("/:id", func(context *gin.Context) {
		postcontroller.GetPost(context, db)
	})

	router.GET("/search", func(context *gin.Context) {
		context.JSON(http.StatusNotImplemented, gin.H{
			"success": false,
			"message": "Rota não implementada",
		})
	})

	router.PUT("", middlewares.AuthenticateUser, func(context *gin.Context) {
		postcontroller.EditPost(context, db)
	})

	router.DELETE("/:id", middlewares.AuthenticateUser, func(context *gin.Context) {
		postcontroller.DeletePost(context, db)
	})
}
