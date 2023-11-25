package timeline

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	postrepository "SociLinkApi/repository/post"
	timelinerepository "SociLinkApi/repository/timeline"
	userrepository "SociLinkApi/repository/user"
	types "SociLinkApi/types/pagination"
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

	uid, exists := context.Get("userId")
	var userId *uuid.UUID
	if exists {
		id := uid.(uuid.UUID)
		userId = &id
	}

	var err error
	posts := types.PostListing{
		PaginationResponse: types.PaginationResponse{
			Page:     pagination.Page,
			PageSize: pagination.PageSize,
		},
	}

	if userId != nil && *userId == user.ID {
		err = postrepository.GetPostsByUserId(*userId, &posts, db)
	} else {
		err = timelinerepository.GetUserTimeline(userId, user.ID, &posts, db)
	}

	if err != nil {
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
