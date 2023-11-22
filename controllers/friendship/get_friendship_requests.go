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

	var pagination dto.PaginationRequestDto
	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if requests, err := frienshiprepository.GetFriendshipRequests(uid.(uuid.UUID), pagination, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetFriendshipRequestsResponseDto{
			PaginationResponse: requests.PaginationResponse,
			Requests:           make([]dto.FriendshipResponseDto, len(requests.Friendships)),
		}

		for i, request := range requests.Friendships {
			request.Friend = request.User

			response.Requests[i] = dto.FriendshipToFriendshipResponseDto(request)
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    response,
		})
	}
}
