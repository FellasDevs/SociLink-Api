package postcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	likerepository "SociLinkApi/repository/like"
	postrepository "SociLinkApi/repository/post"
	authtypes "SociLinkApi/types/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func EditPost(context *gin.Context, db *gorm.DB) {
	var postData dto.EditPostRequestDto
	if err := context.ShouldBindJSON(&postData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	postId, err := uuid.Parse(postData.Id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	post := models.Post{ID: postId}
	if err = postrepository.GetPost(&post, &userId, db); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if postData.Images != nil {
		post.Images = postData.Images
	}
	if postData.Content != "" {
		post.Content = postData.Content
	}
	if postData.Visibility != "" {
		if postData.Visibility != "public" && postData.Visibility != "private" && postData.Visibility != "friends" {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Visibilidade deve ser public, private ou friends.",
			})
			return
		}

		visibility := authtypes.Public

		if postData.Visibility == "private" {
			visibility = authtypes.Private
		} else if visibility == "friends" {
			visibility = authtypes.Friends
		}

		post.Visibility = string(visibility)
	}

	if err = postrepository.UpdatePost(&post, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	likes, _ := likerepository.GetPostLikes(post.ID, db)

	userLikedPost := false
	for _, like := range likes {
		if like.UserID == userId {
			userLikedPost = true
			break
		}
	}

	response := dto.CreatePostResponseDto{
		Post: dto.PostToPostResponseDto(post, len(likes), userLikedPost),
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post editado com sucesso!",
		"data":    response,
	})
}
