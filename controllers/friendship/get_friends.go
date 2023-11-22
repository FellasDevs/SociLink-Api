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

	var pagination dto.PaginationRequestDto
	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if friendships, err := frienshiprepository.GetFriendships(userId, pagination, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetFriendsResponseDto{
			PaginationResponse: friendships.PaginationResponse,
			Friends:            make([]dto.FriendshipResponseDto, len(friendships.Friendships)),
		}

		for i, friendship := range friendships.Friendships {
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
