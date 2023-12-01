package friendshipcontroller

import (
	"SociLinkApi/dto"
	friendshiprepository "SociLinkApi/repository/friendship"
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

	if requests, err := friendshiprepository.GetFriendshipRequests(uid.(uuid.UUID), pagination, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetFriendshipRequestsResponseDto{
			Requests: make([]dto.FriendshipResponseDto, len(requests)),
		}

		for i, request := range requests {
			request.Friend = request.User

			response.Requests[i] = dto.FriendshipToResponseDto(request)
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    response,
		})
	}
}
