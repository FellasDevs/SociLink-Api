package usercontroller

import (
	"SociLinkApi/dto"
	userrepository "SociLinkApi/repository/user"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetSelf(context *gin.Context, db *gorm.DB) {
	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "userId not found in context",
		})
		return
	}

	if user, err := userrepository.GetUserById(userId.(uuid.UUID), db); err != nil {
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
	} else {
		response := dto.GetUserByIdResponseDto{User: dto.UserResponseDto{
			Name:      user.Name,
			Birthdate: user.Birthdate.String(),
		}}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "usuário obtido com sucesso",
			"data":    response,
		})
	}
}
