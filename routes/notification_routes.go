package routes

import (
	"SociLinkApi/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NotificationRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", middlewares.AuthenticateUser, func(context *gin.Context) {

	})
}
