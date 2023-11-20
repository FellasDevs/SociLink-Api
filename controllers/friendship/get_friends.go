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
			user := friendship.User
			if friendship.UserID == userId {
				user = friendship.Friend
			}

			response.Friends[i] = dto.FriendshipResponseDto{
				Id: friendship.ID.String(),
				Friend: dto.UserResponseDto{
					Id:        user.ID.String(),
					Name:      user.Name,
					Birthdate: user.Birthdate.String(),
					Nickname:  user.Nickname,
					Country:   user.Country,
					City:      user.City,
					Picture:   user.Picture,
					Banner:    user.Banner,
					CreatedAt: user.CreatedAt.String(),
				},
				CreatedAt: friendship.CreatedAt,
			}
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    response,
		})
	}
}
