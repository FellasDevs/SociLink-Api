package routes

import (
	commentcontroller "SociLinkApi/controllers/comment"
	"SociLinkApi/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CommentRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/", middlewares.AuthenticateUserOptionally, func(context *gin.Context) {
		commentcontroller.GetPostComments(context, db)
	})

	router.POST("/", middlewares.AuthenticateUser, func(context *gin.Context) {
		commentcontroller.CreateComment(context, db)
	})

	router.PUT("/:id", middlewares.AuthenticateUser, func(context *gin.Context) {
		commentcontroller.EditComment(context, db)
	})

	router.DELETE("/:id", middlewares.AuthenticateUser, func(context *gin.Context) {
		commentcontroller.DeleteComment(context, db)
	})
	router.POST("/", middlewares.AuthenticateUser, func(context *gin.Context) {
		commentcontroller.CreateCommentLike(context, db)
	})
	router.DELETE("/", middlewares.AuthenticateUser, func(context *gin.Context) {
		commentcontroller.DeleteCommentLike(context, db)
	})
	router.POST("/:id/replies", middlewares.AuthenticateUser, func(context *gin.Context) {
		commentcontroller.CreateCommentReply(context, db)
	})

	router.DELETE("/:id/replies", middlewares.AuthenticateUser, func(context *gin.Context) {
		commentcontroller.DeleteCommentReply(context, db)
	})

}
