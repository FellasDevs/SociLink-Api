package usercontroller

import (
	"SociLinkApi/dto"
	userrepository "SociLinkApi/repository/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func SearchUsers(context *gin.Context, db *gorm.DB) {
	var params dto.SearchUsersRequestDto
	if err := context.ShouldBindQuery(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if users, err := userrepository.GetUsersByNameOrNickname(params.Search, params.PaginationRequestDto, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.SearchUsersResponseDto{
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
