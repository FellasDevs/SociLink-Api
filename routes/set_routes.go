package routes

import (
	authRoutes "SociLinkApi/routes/auth"
	userRoutes "SociLinkApi/routes/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutes(router *gin.Engine, db *gorm.DB) {
	userRoutes.UserRoutes(router.Group("/users"), db)
	authRoutes.AuthRoutes(router.Group("/auth"), db)
}
