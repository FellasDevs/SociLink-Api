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

func CreatePost(context *gin.Context, db *gorm.DB) {
	var postData dto.CreatePostRequestDto

	if err := context.ShouldBindJSON(&postData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	uid, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "token de autenticação não encontrado",
		})
		return
	}

	visibility := authtypes.Public

	if postData.Visibility == "private" {
		visibility = authtypes.Private
	} else if postData.Visibility == "friends" {
		visibility = authtypes.Friends
	}

	post := models.Post{
		UserID:     uid.(uuid.UUID),
		Content:    postData.Content,
		Images:     postData.Images,
		Visibility: string(visibility),
	}

	err := postrepository.CreatePost(&post, db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Post criado com sucesso!",
	})
}
