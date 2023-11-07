package postcontroller

import (
	"SociLinkApi/dto"
	postrepository "SociLinkApi/repository/post"
	"errors"
	"net/http"

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

	post, err := postrepository.GetPost(postId, db)

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
	} else {
		response := dto.GetPostResponseDto{
			Id: post.ID.String(),
			User: dto.PayloadUser{
				Name:      post.User.Name,
				Birthdate: post.User.Birthdate.String(),
			},
			Content:    post.Content,
			Visibility: post.Visibility,
			Images:     post.Images,
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Post obtido com sucesso",
			"data":    response,
		})
	}
}
