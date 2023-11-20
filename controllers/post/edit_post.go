package postcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
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

	post := models.Post{ID: postId}
	if err = postrepository.GetPost(&post, db); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	uid, _ := context.Get("userId")

	userId := uid.(uuid.UUID)
	if userId != post.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "você não tem permissão para editar esse post",
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

	err = postrepository.UpdatePost(&post, db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response := dto.CreatePostResponseDto{
		Post: dto.PostResponseDto{
			Id: post.ID.String(),
			User: dto.UserResponseDto{
				Id:        post.User.ID.String(),
				Name:      post.User.Name,
				Nickname:  post.User.Nickname,
				Birthdate: post.User.Birthdate.String(),
				Country:   post.User.Country,
				City:      post.User.City,
				Picture:   post.User.Picture,
				Banner:    post.User.Banner,
				CreatedAt: post.User.CreatedAt.String(),
			},
			Content:    post.Content,
			Images:     post.Images,
			Visibility: post.Visibility,
		},
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post editado com sucesso!",
		"data":    response,
	})
}
