package routes

import (
	routes "SociLinkApi/routes/user"
	"SociLinkApi/routes/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(router *gin.Engine, db *gorm.DB) {
	routes.UserRoutes(router.Group("/users"), db)
	authroutes.AuthRouterGroup(router.Group("/auth"), db)
}
