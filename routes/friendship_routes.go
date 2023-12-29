package routes

import (
	"SociLinkApi/controllers/friendship"
	"SociLinkApi/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FriendshipRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("", func(context *gin.Context) {
		friendshipcontroller.GetFriends(context, db)
	})

	router.GET("/:nickname", middlewares.AuthenticateUser, func(context *gin.Context) {
		friendshipcontroller.GetFriendship(context, db)
	})

	router.GET("/requests", middlewares.AuthenticateUser, func(context *gin.Context) {
		friendshipcontroller.GetAllFriendshipRequests(context, db)
	})

	router.POST("/requests/:id", middlewares.AuthenticateUser, func(context *gin.Context) {
		friendshipcontroller.RequestFriendship(context, db)
	})

	router.PUT("/requests", middlewares.AuthenticateUser, func(context *gin.Context) {
		friendshipcontroller.AnswerFriendshipRequest(context, db)
	})

	router.DELETE("/:id", middlewares.AuthenticateUser, func(context *gin.Context) {
		friendshipcontroller.RemoveFriend(context, db)
	})
}
