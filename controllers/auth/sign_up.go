package authcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	userrepository "SociLinkApi/repository/user"
	authservice "SociLinkApi/services/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func SignUpController(context *gin.Context, db *gorm.DB) {
	var userInfo dto.SignUpRequestDto

	if err := context.ShouldBindJSON(&userInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	password, err := authservice.EncryptPassword(userInfo.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	birthdate, err := authservice.ParseBirthdate(userInfo.Birthdate)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	user := models.User{
		Name:      userInfo.Name,
		Email:     userInfo.Email,
		Password:  password,
		Birthdate: birthdate,
	}

	createdUser, err := userrepository.CreateUser(user, db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	token, err := authservice.CreateJWT(createdUser.Id, time.Hour*24)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	response := dto.SignUpResponseDto{
		User: dto.PayloadUser{
			Name:      user.Name,
			Email:     user.Email,
			Birthdate: user.Birthdate.String(),
		},
		AuthToken: token,
	}

	context.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "Usuário criado com sucesso!",
		"response": response,
	})
}
