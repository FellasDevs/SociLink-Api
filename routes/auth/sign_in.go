package authroutes

import (
	authcontroller "SociLinkApi/controllers/auth"
	"SociLinkApi/dto"
	"SociLinkApi/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SignIn(context *gin.Context, db *gorm.DB) {
	var userInfo dto.SignInDto

	if err := context.ShouldBindJSON(&userInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	_, token, err := authcontroller.SignInController(userInfo, db)
	if err != nil {
		utils.SendUnknownError(err, context)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success":   true,
		"message":   "Usuário logado com sucesso!",
		"authToken": token,
	})
}
