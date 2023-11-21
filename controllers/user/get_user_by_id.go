package usercontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
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

	user := models.User{ID: id}
	if err = userrepository.GetUser(&user, db); err != nil {
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

	response := dto.GetUserByIdResponseDto{
		User: dto.UserToUserResponseDto(user),
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "usuário obtido com sucesso",
		"data":    response,
	})
}
