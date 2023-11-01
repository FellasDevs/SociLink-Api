package userroutes

import (
	usercontroller "SociLinkApi/controllers/user"
	"SociLinkApi/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SearchUsersRoute(context *gin.Context, db *gorm.DB) {
	search := context.Param("search")

	if search == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Search query cannot be empty",
		})
		return
	}

	user, err := usercontroller.SearchUsers(search, db)

	if err != nil {
		utils.SendUnknownError(err, context)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"user":    user,
	})
}
