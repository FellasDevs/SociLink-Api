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

	if friends, err := frienshiprepository.GetFriendships(uid.(uuid.UUID), db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := make([]dto.FriendshipResponseDto, len(friends))

		for i, friend := range friends {
			response[i] = dto.FriendshipResponseDto{
				Id: friend.ID,
				User: dto.UserResponseDto{
					Id:        friend.User.ID.String(),
					Name:      friend.User.Name,
					Birthdate: friend.User.Birthdate.String(),
				},
				CreatedAt: friend.CreatedAt,
			}
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    response,
		})
	}
}
