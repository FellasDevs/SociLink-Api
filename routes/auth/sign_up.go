package authroutes

import (
	authcontroller "SociLinkApi/controllers/auth"
	"SociLinkApi/dto"
	"SociLinkApi/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SignUp(context *gin.Context, db *gorm.DB) {
	var userInfo dto.SignUpDto

	if err := context.ShouldBindJSON(&userInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	_, token, err := authcontroller.SignUpController(userInfo, db)
	if err != nil {
		utils.SendUnknownError(err, context)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success":   true,
		"message":   "Usuário criado com sucesso!",
		"authToken": token,
	})

}
