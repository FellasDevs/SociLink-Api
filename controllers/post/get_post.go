package postcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	likerepository "SociLinkApi/repository/like"
	postrepository "SociLinkApi/repository/post"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetPost(context *gin.Context, db *gorm.DB) {
	idString := context.Param("id")
	postId, err := uuid.Parse(idString)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "O id informado é inválido",
		})
		return
	}

	var userId *uuid.UUID
	uid, exists := context.Get("userId")
	if exists {
		id := uid.(uuid.UUID)
		userId = &id
	}

	post := models.Post{ID: postId}
	if err = postrepository.GetPost(&post, userId, db); err != nil {
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

	response := dto.GetPostResponseDto{
		Post: dto.PostToPostResponseDto(post, len(likes), userLikedPost),
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post obtido com sucesso",
		"data":    response,
	})
}
