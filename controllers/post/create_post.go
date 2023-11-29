package postcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	postrepository "SociLinkApi/repository/post"
	authtypes "SociLinkApi/types/auth"
	"net/http"
	"strings"

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

	var fieldErrors []string

	if postData.Content == "" {
		fieldErrors = append(fieldErrors, "Conteúdo não pode ser vazio.")
	}
	if postData.Visibility != "public" && postData.Visibility != "private" && postData.Visibility != "friends" {
		fieldErrors = append(fieldErrors, "Visibilidade deve ser public, private ou friends.")
	}

	if len(fieldErrors) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": strings.Join(fieldErrors, " "),
		})
		return
	}

	visibility := authtypes.Public

	if postData.Visibility == "private" {
		visibility = authtypes.Private
	} else if postData.Visibility == "friends" {
		visibility = authtypes.Friends
	}

	uid, _ := context.Get("userId")

	post := models.Post{
		UserID:     uid.(uuid.UUID),
		Content:    postData.Content,
		Images:     postData.Images,
		Visibility: string(visibility),
	}

	if err := postrepository.CreatePost(&post, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	response := dto.CreatePostResponseDto{
		Post: dto.PostToPostResponseDto(post, 0),
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Post criado com sucesso!",
		"data":    response,
	})
}
