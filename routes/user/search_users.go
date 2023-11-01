package userroutes

import (
	usercontroller "SociLinkApi/controllers/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SearchUsersRoute(context *gin.Context, db *gorm.DB) {
	search := context.Param("search")

	users, err := usercontroller.GetUsersByName(search, db)

	if err != nil {
		context.JSON(err.StatusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"users":   users,
	})
}
