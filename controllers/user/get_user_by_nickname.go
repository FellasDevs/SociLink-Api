package usercontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	userrepository "SociLinkApi/repository/user"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetUserByNickname(context *gin.Context, db *gorm.DB) {
	nickname := context.Param("nick")
	if nickname == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Apelido não informado",
		})
		return
	}

	user := models.User{Nickname: nickname}
	if err := userrepository.GetUser(&user, db); err != nil {
		var statusCode int

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}

		context.JSON(statusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response := dto.GetUserResponseDto{User: dto.UserResponseDto{
		Id:        user.ID.String(),
		Name:      user.Name,
		Nickname:  user.Nickname,
		Birthdate: user.Birthdate.String(),
		Country:   user.Country,
		City:      user.City,
		Picture:   user.Picture,
		Banner:    user.Banner,
		CreatedAt: user.CreatedAt.String(),
	}}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "usuário obtido com sucesso",
		"data":    response,
	})
}
