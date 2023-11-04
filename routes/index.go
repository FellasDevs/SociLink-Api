package routes

import (
	"SociLinkApi/routes/auth"
	userroutes "SociLinkApi/routes/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(router *gin.Engine, db *gorm.DB) {
	userroutes.UserRouterGroup(router.Group("/users"), db)

	authroutes.AuthRouterGroup(router.Group("/auth"), db)
}
