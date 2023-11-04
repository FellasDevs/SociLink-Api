package authroutes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouterGroup(router *gin.RouterGroup, db *gorm.DB) {
	router.POST("/sign_in", func(context *gin.Context) {
		SignIn(context, db)
	})

	router.POST("/sign_up", func(context *gin.Context) {
		SignUp(context, db)
	})
}
