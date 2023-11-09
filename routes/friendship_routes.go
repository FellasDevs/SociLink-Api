package routes

import (
	"SociLinkApi/controllers/friendship"
	"SociLinkApi/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FriendshipRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.Use(middlewares.AuthenticateUser)

	router.GET("", func(context *gin.Context) {
		friendshipcontroller.GetFriends(context, db)
	})

	router.GET("/requests", func(context *gin.Context) {
		friendshipcontroller.GetAllFriendshipRequests(context, db)
	})

	router.POST("/:id", func(context *gin.Context) {
		friendshipcontroller.RequestFriendship(context, db)
	})

	router.POST("/answer", func(context *gin.Context) {
		friendshipcontroller.AnswerFriendshipRequests(context, db)
	})

	router.DELETE("/:id", func(context *gin.Context) {
		friendshipcontroller.RemoveFriend(context, db)
	})
}
