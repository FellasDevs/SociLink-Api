package friendshipcontroller

import (
	"SociLinkApi/dto"
	frienshiprepository "SociLinkApi/repository/frienship"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetAllFriendshipRequests(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")

	if requests, err := frienshiprepository.GetFriendshipRequests(uid.(uuid.UUID), db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetFriendshipRequestsResponseDto{
			Requests: make([]dto.FriendshipResponseDto, len(requests)),
		}

		for i, request := range requests {
			response.Requests[i] = dto.FriendshipResponseDto{
				Id: request.ID,
				User: dto.UserResponseDto{
					Id:        request.User.ID.String(),
					Name:      request.User.Name,
					Birthdate: request.User.Birthdate.String(),
					Nickname:  request.User.Nickname,
					Country:   request.User.Country,
					City:      request.User.City,
					Picture:   request.User.Picture,
					Banner:    request.User.Banner,
					CreatedAt: request.User.CreatedAt.String(),
				},
				CreatedAt: request.CreatedAt,
			}
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    response,
		})
	}
}
