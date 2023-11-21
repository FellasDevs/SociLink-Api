package friendshipcontroller

import (
	"SociLinkApi/dto"
	frienshiprepository "SociLinkApi/repository/frienship"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetFriends(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	if friendships, err := frienshiprepository.GetFriendships(userId, db); err != nil {
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

			response.Friends[i] = dto.FriendshipToFriendshipResponseDto(friendship)
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    response,
		})
	}
}
