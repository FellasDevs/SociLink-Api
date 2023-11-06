package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func PostRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/create", func(context *gin.Context) {
		context.JSON(http.StatusNotImplemented, gin.H{
			"success": false,
			"message": "Rota não implementada",
		})
	})

	router.GET("/:id", func(context *gin.Context) {
		context.JSON(http.StatusNotImplemented, gin.H{
			"success": false,
			"message": "Rota não implementada",
		})
	})

	router.GET("/search", func(context *gin.Context) {
		context.JSON(http.StatusNotImplemented, gin.H{
			"success": false,
			"message": "Rota não implementada",
		})
	})

	router.PUT("/:id", func(context *gin.Context) {
		context.JSON(http.StatusNotImplemented, gin.H{
			"success": false,
			"message": "Rota não implementada",
		})
	})

	router.DELETE("/:id", func(context *gin.Context) {
		context.JSON(http.StatusNotImplemented, gin.H{
			"success": false,
			"message": "Rota não implementada",
		})
	})
}
