package timeline

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	likerepository "SociLinkApi/repository/like"
	postrepository "SociLinkApi/repository/post"
	timelinerepository "SociLinkApi/repository/timeline"
	userrepository "SociLinkApi/repository/user"
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
	var posts []models.Post

	if userId != nil && *userId == user.ID {
		posts, err = postrepository.GetPostsByUserId(*userId, pagination, db)
	} else {
		posts, err = timelinerepository.GetUserTimeline(userId, user.ID, pagination, db)
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetUserTimelineResponseDto{
			User:  dto.UserToUserWithFriendsResponseDto(user),
			Posts: make([]dto.PostResponseDto, len(posts)),
		}

		for i, post := range posts {
			likes, _ := likerepository.GetPostLikes(post.ID, db)

			userLikedPost := false
			if userId != nil {
				for _, like := range likes {
					if like.UserID == *userId {
						userLikedPost = true
						break
					}
				}
			}

			response.Posts[i] = dto.PostToPostResponseDto(post, len(likes), userLikedPost)
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "posts recuperados com sucesso",
			"data":    response,
		})
	}
}
