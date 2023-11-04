package routes

import (
	authcontroller "SociLinkApi/controllers/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/sign_in", func(context *gin.Context) {
		authcontroller.SignInController(context, db)
	})

	router.POST("/sign_up", func(context *gin.Context) {
		authcontroller.SignUpController(context, db)
	})
}
