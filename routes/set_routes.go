package routes

import (
	userroutes "SociLinkApi/routes/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(router *gin.Engine, db *gorm.DB) {
	userroutes.SetUserRoutes(router.Group("/users"), db)
}
