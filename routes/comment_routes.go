package routes

import (
	"SociLinkApi/controllers/comment"
	"SociLinkApi/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("", middlewares.AuthenticateUserOptionally, func(context *gin.Context) {
		commentcontroller.GetPostComments(context, db)
	})

	router.POST("", middlewares.AuthenticateUser, func(context *gin.Context) {
		commentcontroller.CreateComment(context, db)
	})

	router.PUT("/:id", middlewares.AuthenticateUser, func(context *gin.Context) {
		commentcontroller.EditComment(context, db)
	})

	router.DELETE("/:id", middlewares.AuthenticateUser, func(context *gin.Context) {
		commentcontroller.DeleteComment(context, db)
	})
}
