package userroutes

import (
	usercontroller "SociLinkApi/controllers/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetUserRoute(context *gin.Context, db *gorm.DB) {
	id := context.Param("id")

	user, err := usercontroller.GetUserById(id, db)

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
		"user":    user,
	})
}
