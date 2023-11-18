package routes

import (
	usercontroller "SociLinkApi/controllers/user"
	"SociLinkApi/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/self", middlewares.AuthenticateUser, func(context *gin.Context) {
		usercontroller.GetSelf(context, db)
	})

	router.PUT("/self", middlewares.AuthenticateUser, func(context *gin.Context) {
		usercontroller.EditUserInfo(context, db)
	})

	router.GET("/:id", func(context *gin.Context) {
		usercontroller.GetUserById(context, db)
	})

	router.GET("/search/:search", func(context *gin.Context) {
		usercontroller.GetUsersByName(context, db)
	})
}
