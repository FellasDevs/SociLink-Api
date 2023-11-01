package userroutes

import (
	usercontroller "SociLinkApi/controllers/user"
	"SociLinkApi/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetUserRoute(context *gin.Context, db *gorm.DB) {
	idString := context.Param("id")

	id, err := uuid.Parse(idString)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user, err := usercontroller.GetUser(id, db)

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
