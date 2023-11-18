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

	if users, err := userrepository.GetUsersByNameOrNickname(search, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetUsersByNameResponseDto{Users: make([]dto.UserResponseDto, len(users))}

		for i, user := range users {
			response.Users[i] = dto.UserResponseDto{
				Id:        user.ID.String(),
				Name:      user.Name,
				Birthdate: user.Birthdate.String(),
				Nickname:  user.Nickname,
				Country:   user.Country,
				City:      user.City,
				Picture:   user.Picture,
				Banner:    user.Banner,
				CreatedAt: user.CreatedAt.String(),
			}
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "usuários obtidos com sucesso",
			"data":    response,
		})
	}
}
