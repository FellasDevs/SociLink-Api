package postcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	frienshiprepository "SociLinkApi/repository/frienship"
	postrepository "SociLinkApi/repository/post"
	authtypes "SociLinkApi/types/auth"
	"errors"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetPost(context *gin.Context, db *gorm.DB) {
	idString := context.Param("id")

	postId, err := uuid.Parse(idString)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id is not a valid uuid",
		})
		return
	}

	post := models.Post{ID: postId}
	if err = postrepository.GetPost(&post, db); err != nil {
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

	var visibility authtypes.Visibility

	uid, exists := context.Get("userId")
	if !exists {
		visibility = authtypes.Public
	} else if uid.(uuid.UUID) == post.User.ID {
		visibility = authtypes.Private
	} else if _, err := frienshiprepository.GetFriendshipByUsers(uid.(uuid.UUID), post.User.ID, db); err == nil {
		visibility = authtypes.Friends
	}

	if !slices.Contains(visibility.GetAllowedVisibilities(), post.Visibility) {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "post not found",
		})
		return
	}

	response := dto.PostToPostResponseDto(post)

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post obtido com sucesso",
		"data":    response,
	})
}
