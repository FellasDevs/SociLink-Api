package friendshipcontroller

import (
	"SociLinkApi/dto"
	friendshiprepository "SociLinkApi/repository/friendship"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetFriends(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	var pagination dto.PaginationRequestDto
	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if friendships, err := friendshiprepository.GetFriendships(userId, pagination, db); err != nil {
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
