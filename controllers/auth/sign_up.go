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
			"message": err.Error(),
		})
		return
	}

	birthdate, err := authservice.ParseBirthdate(userInfo.Birthdate)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	password, err := authservice.EncryptPassword(userInfo.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user := models.User{
		Name:      userInfo.Name,
		Email:     userInfo.Email,
		Password:  password,
		Birthdate: birthdate,
		Nickname:  userInfo.Nickname,
		Country:   userInfo.Country,
		City:      userInfo.City,
	}

	err = userrepository.CreateUser(&user, db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	token, err := authservice.CreateAuthToken(user.ID, time.Hour*24)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response := dto.SignUpResponseDto{
		User: dto.UserResponseDto{
			Id:        user.ID.String(),
			Name:      user.Name,
			Nickname:  user.Nickname,
			Birthdate: user.Birthdate.String(),
			Country:   user.Country,
			City:      user.City,
			CreatedAt: user.CreatedAt.String(),
		},
		AuthToken: token,
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Usuário criado com sucesso!",
		"data":    response,
	})
}
