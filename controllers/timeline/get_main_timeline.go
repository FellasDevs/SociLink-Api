package timeline

import (
	"SociLinkApi/dto"
	frienshiprepository "SociLinkApi/repository/friendship"
	postrepository "SociLinkApi/repository/post"
	authtypes "SociLinkApi/types/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetMainTimeline(context *gin.Context, db *gorm.DB) {
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

	friends, err := frienshiprepository.GetAllFriendships(userId, db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	userIds := make([]uuid.UUID, len(friends)+1)

	userIds[0] = userId
	for i, friend := range friends {
		id := friend.FriendID

		if friend.FriendID == userId {
			id = friend.UserID
		}

		userIds[i+1] = id
	}

	if posts, err := postrepository.GetPostsByUsers(userIds, authtypes.Friends, pagination, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetMainTimelineResponseDto{
			PaginationResponse: posts.PaginationResponse,
			Posts:              make([]dto.PostResponseDto, len(posts.Posts)),
		}

		for i, post := range posts.Posts {
			response.Posts[i] = dto.PostToPostResponseDto(post)
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "posts recuperados com sucesso",
			"data":    response,
		})
	}
}
