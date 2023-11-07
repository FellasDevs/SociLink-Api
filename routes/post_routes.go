package routes

import (
	postcontroller "SociLinkApi/controllers/post"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/create", func(context *gin.Context) {
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

	router.PUT("/:id", func(context *gin.Context) {
		postcontroller.EditPost(context, db)
	})

	router.DELETE("/:id", func(context *gin.Context) {
		postcontroller.DeletePost(context, db)
	})
}
