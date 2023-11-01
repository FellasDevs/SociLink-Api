package userroutes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouterGroup(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/:id", func(context *gin.Context) {
		GetUserRoute(context, db)
	})

	router.GET("/search/:search", func(context *gin.Context) {
		SearchUsersRoute(context, db)
	})
}
