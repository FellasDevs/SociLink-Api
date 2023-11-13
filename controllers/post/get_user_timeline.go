package postcontroller

import (
	"SociLinkApi/dto"
	postrepository "SociLinkApi/repository/post"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetUserTimeline(context *gin.Context, db *gorm.DB) {
	uid := context.Param("id")

	userId, err := uuid.Parse(uid)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if posts, err := postrepository.GetPostsByUser(userId, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		var response dto.TimelineResponseDto

		response.Posts = make([]dto.GetPostResponseDto, len(posts))

		for i, post := range posts {
			response.Posts[i] = dto.GetPostResponseDto{
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
			"message": "posts recuperados com sucesso",
			"data":    response.Posts,
		})
	}
}
