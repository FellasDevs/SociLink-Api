package routes

import (
	usercontroller "SociLinkApi/controllers/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/:id", func(context *gin.Context) {
		usercontroller.GetUserById(context, db)
	})

	router.GET("/search/:search", func(context *gin.Context) {
		usercontroller.GetUsersByName(context, db)
	})
}
