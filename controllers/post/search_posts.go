package postcontroller

import (
	"SociLinkApi/dto"
	postrepository "SociLinkApi/repository/post"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func SearchPosts(context *gin.Context, db *gorm.DB) {
	search := context.Param("search")
	if search == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "pesquisa inválida",
		})
		return
	}

	uid, exist := context.Get("userId")
	var userId *uuid.UUID
	if exist {
		id := uid.(uuid.UUID)
		userId = &id
	}

	if posts, err := postrepository.SearchPosts(search, userId, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.SearchPostResponseDto{
			Posts: make([]dto.PostResponseDto, len(posts)),
		}

		for i, post := range posts {
			response.Posts[i] = dto.PostResponseDto{
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
			}
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "posts encontrados",
			"data":    response,
		})
	}
}
