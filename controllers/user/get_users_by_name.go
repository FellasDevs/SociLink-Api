package usercontroller

import (
	"SociLinkApi/dto"
	userrepository "SociLinkApi/repository/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetUsersByName(context *gin.Context, db *gorm.DB) {
	search := context.Param("search")

	if search == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "A pesquisa não pode ser vazia",
		})
		return
	}

	var pagination dto.PaginationRequestDto
	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if users, err := userrepository.GetUsersByNameOrNickname(search, pagination, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetUsersByNameResponseDto{
			PaginationResponse: users.PaginationResponse,
			Users:              make([]dto.UserResponseDto, len(users.Users)),
		}

		for i, user := range users.Users {
			response.Users[i] = dto.UserToUserResponseDto(user)
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "usuários obtidos com sucesso",
			"data":    response,
		})
	}
}
