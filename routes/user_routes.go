package routes

import (
	usercontroller "SociLinkApi/controllers/user"
	"SociLinkApi/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.Use(func(context *gin.Context) {
		middlewares.AuthenticateUser(context, db)
	})

	router.GET("/:id", func(context *gin.Context) {
		usercontroller.GetUserById(context, db)
	})

	router.GET("/search/:search", func(context *gin.Context) {
		usercontroller.GetUsersByName(context, db)
	})
}
