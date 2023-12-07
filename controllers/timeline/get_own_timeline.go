package timeline

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	likerepository "SociLinkApi/repository/like"
	postrepository "SociLinkApi/repository/post"
	timelinerepository "SociLinkApi/repository/timeline"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetOwnTimeline(context *gin.Context, db *gorm.DB) {
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

	if posts, err := timelinerepository.GetOwnTimeline(userId, pagination, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetMainTimelineResponseDto{
			Posts: make([]dto.PostResponseDto, len(posts)),
		}

		for i, post := range posts {
			likes, _ := likerepository.GetPostLikes(post.ID, db)

			userLikedPost := false
			for _, like := range likes {
				if like.UserID == userId {
					userLikedPost = true
					break
				}
			}

			response.Posts[i] = dto.PostToResponseDto(post, len(likes), userLikedPost)

			if post.OriginalPostID != nil {
				originalPost := models.Post{
					ID: *post.OriginalPostID,
				}

				err = postrepository.GetPost(&originalPost, &userId, db)

				if err == nil {
					likes, _ = likerepository.GetPostLikes(post.ID, db)

					userLikedPost = false
					for _, like := range likes {
						if like.UserID == userId {
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
