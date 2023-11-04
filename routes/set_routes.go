package routes

import (
	"SociLinkApi/routes/auth"
	routes "SociLinkApi/routes/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(router *gin.Engine, db *gorm.DB) {
	routes.UserRoutes(router.Group("/users"), db)
	routes.AuthRoutes(router.Group("/auth"), db)
}
