package authcontroller

import (
	"SociLinkApi/dto"
	userrepository "SociLinkApi/repository/user"
	authservice "SociLinkApi/services/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

func SignInController(context *gin.Context, db *gorm.DB) {
	var userInfo dto.SignInRequestDto

	if err := context.ShouldBindJSON(&userInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	var fieldErrors []string

	if userInfo.Email == "" {
		fieldErrors = append(fieldErrors, "Email não informado.")
	}
	if userInfo.Password == "" {
		fieldErrors = append(fieldErrors, "Senha não informada.")
	}

	if len(fieldErrors) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": strings.Join(fieldErrors, " "),
		})
		return
	}

	user, err := userrepository.GetUserByEmail(userInfo.Email, db)

	if err != nil {
		var statusCode int
		message := err.Error()

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			message = "Email não cadastrado."
		} else {
			statusCode = http.StatusInternalServerError
		}

		context.JSON(statusCode, gin.H{
			"success": false,
			"message": message,
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInfo.Password)); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Senha incorreta.",
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

	response := dto.SignInResponseDto{
		User: dto.UserResponseDto{
			Id:        user.ID.String(),
			Name:      user.Name,
			Nickname:  user.Nickname,
			Birthdate: user.Birthdate.String(),
			Country:   user.Country,
			City:      user.City,
			Picture:   user.Picture,
			Banner:    user.Banner,
			CreatedAt: user.CreatedAt.String(),
		},
		AuthToken: token,
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Usuário logado com sucesso!",
		"data":    response,
	})
}
