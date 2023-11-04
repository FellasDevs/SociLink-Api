package usercontroller

import (
	userrepository "SociLinkApi/repository/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetUsersByName(context *gin.Context, db *gorm.DB) {
	search := context.Param("search")

	if search == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "search cannot be empty",
		})
		return
	}

	if users, err := userrepository.GetUsersByName(search, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "usuários obtidos com sucesso",
			"users":   users,
		})
	}
}
