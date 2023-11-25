package routes

import (
	"SociLinkApi/controllers/timeline"
	"SociLinkApi/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TimelineRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("", middlewares.AuthenticateUser, func(context *gin.Context) {
		timeline.GetOwnTimeline(context, db)
	})

	router.GET("/:nick", middlewares.AuthenticateUserOptionally, func(context *gin.Context) {
		timeline.GetUserTimeline(context, db)
	})
}
