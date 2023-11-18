package postcontroller

import (
	"SociLinkApi/dto"
	frienshiprepository "SociLinkApi/repository/frienship"
	postrepository "SociLinkApi/repository/post"
	authtypes "SociLinkApi/types/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetUserTimeline(context *gin.Context, db *gorm.DB) {
	paramId := context.Param("id")

	paramUserId, err := uuid.Parse(paramId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	visibility := authtypes.Public

	uid, exists := context.Get("userId")
	if exists {
		userId := uid.(uuid.UUID)

		if userId == paramUserId {
			visibility = authtypes.Private
		} else if _, err := frienshiprepository.GetFriendshipByUsers(userId, paramUserId, db); err == nil {
			visibility = authtypes.Friends
		}
	}

	if posts, err := postrepository.GetPostsByUser(paramUserId, visibility, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := make([]dto.PostResponseDto, len(posts))

		for i, post := range posts {
			response[i] = dto.PostResponseDto{
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
			"data":    response,
		})
	}
}
