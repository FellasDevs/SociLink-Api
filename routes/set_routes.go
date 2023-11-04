package routes

import (
	routes "SociLinkApi/routes/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(router *gin.Engine, db *gorm.DB) {
	routes.UserRoutes(router.Group("/users"), db)
}
