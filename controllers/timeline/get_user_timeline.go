package timeline

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	frienshiprepository "SociLinkApi/repository/friendship"
	postrepository "SociLinkApi/repository/post"
	userrepository "SociLinkApi/repository/user"
	authtypes "SociLinkApi/types/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetUserTimeline(context *gin.Context, db *gorm.DB) {
	nickname := context.Param("nick")

	var pagination dto.PaginationRequestDto
	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user := models.User{Nickname: nickname}
	if err := userrepository.GetUserWithFriends(&user, db); err != nil {
		var statusCode int

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}

		context.JSON(statusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	visibility := authtypes.Public

	uid, exists := context.Get("userId")
	if exists {
		userId := uid.(uuid.UUID)

		if userId == user.ID {
			visibility = authtypes.Private
		} else if _, err := frienshiprepository.GetFriendshipByUsers(userId, user.ID, db); err == nil {
			visibility = authtypes.Friends
		}
	}

	if posts, err := postrepository.GetPostsByUser(user.ID, visibility, pagination, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetUserTimelineResponseDto{
			PaginationResponse: posts.PaginationResponse,
			User:               dto.UserToUserWithFriendsResponseDto(user),
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
