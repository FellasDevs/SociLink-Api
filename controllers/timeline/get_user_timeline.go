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

	if nickname == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "O apelido não pode ser vazio",
		})
		return
	}

	var pagination dto.PaginationRequestDto
	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user, err := userrepository.GetUserWithFriendsCount(models.User{Nickname: nickname}, db)
	if err != nil {
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

	var posts []models.Post
	err = nil

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
			User:  dto.UserToUserWithFriendsCountResponseDto(user),
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

			response.Posts[i] = dto.PostToResponseDto(post, len(likes), userLikedPost)

			if post.OriginalPostID != nil {
				originalPost := models.Post{
					ID: *post.OriginalPostID,
				}

				err = postrepository.GetPost(&originalPost, userId, db)

				if err == nil {
					likes, _ = likerepository.GetPostLikes(post.ID, db)

					userLikedPost = false
					for _, like := range likes {
						if like.UserID == *userId {
							userLikedPost = true
							break
						}
					}

					originalPostResponseDto := dto.PostToResponseDto(originalPost, len(likes), userLikedPost)
					response.Posts[i].OriginalPost = &originalPostResponseDto
				}
			}
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "posts recuperados com sucesso",
			"data":    response,
		})
	}
}
