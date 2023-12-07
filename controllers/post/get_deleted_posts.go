package postcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	postrepository "SociLinkApi/repository/post"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetDeletedPosts(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	postModel := models.Post{
		Deleted: true,
		UserID:  userId,
	}

	if posts, err := postrepository.GetPosts(postModel, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"succes":  false,
			"message": err.Error(),
		})
	} else {
		response := dto.GetDeletedPostsResponseDto{
			Posts: make([]dto.PostResponseDto, len(posts)),
		}

		for i, post := range posts {
			response.Posts[i] = dto.PostToResponseDto(post, 0, false)
		}

		context.JSON(200, gin.H{
			"success": true,
			"message": "Postagens recuperadas com sucesso",
			"data":    response,
		})
	}
}
