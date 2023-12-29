package friendshipcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	friendshiprepository "SociLinkApi/repository/friendship"
	userrepository "SociLinkApi/repository/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetFriends(context *gin.Context, db *gorm.DB) {
	var params dto.GetFriendsRequestDto
	if err := context.ShouldBindQuery(&params); err != nil || params.Nickname == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "O apelido do usuário deve ser informado e ser válido",
		})
		return
	}

	user := models.User{Nickname: params.Nickname}

	if err := userrepository.GetUser(&user, db); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if friendships, err := friendshiprepository.GetFriendships(user.ID, params.PaginationRequestDto, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetFriendsResponseDto{
			Friends: make([]dto.FriendshipResponseDto, len(friendships)),
		}

		for i, friendship := range friendships {
			if friendship.FriendID == user.ID {
				friendship.Friend = friendship.User
			}

			response.Friends[i] = dto.FriendshipToResponseDto(friendship)
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    response,
		})
	}
}
