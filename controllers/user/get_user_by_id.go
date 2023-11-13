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

func GetUserById(context *gin.Context, db *gorm.DB) {
	idString := context.Param("id")

	id, err := uuid.Parse(idString)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id is not a valid uuid",
		})
		return
	}

	if user, err := userrepository.GetUserById(id, db); err != nil {
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
			Id:        user.ID.String(),
			Name:      user.Name,
			Birthdate: user.Birthdate.String(),
			Nickname:  user.Nickname,
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
}
