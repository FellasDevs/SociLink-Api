package friendshipcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	friendshiprepository "SociLinkApi/repository/friendship"
	userrepository "SociLinkApi/repository/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetFriends(context *gin.Context, db *gorm.DB) {
	var userId uuid.UUID

	uid, uidExists := context.Get("userId")

	var params dto.GetFriendsRequestDto
	if err := context.ShouldBindQuery(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if uidExists {
		userId = uid.(uuid.UUID)
	} else if params.Nickname != "" {
		user := models.User{Nickname: params.Nickname}

		if err := userrepository.GetUser(&user, db); err != nil {
			context.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		userId = user.ID
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Você deve estar logado ou passar um nickname para acesar essa rota",
		})
		return
	}

	if friendships, err := friendshiprepository.GetFriendships(userId, params.PaginationRequestDto, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetFriendsResponseDto{
			Friends: make([]dto.FriendshipResponseDto, len(friendships)),
		}

		for i, friendship := range friendships {
			if friendship.FriendID == userId {
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
