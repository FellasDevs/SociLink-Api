package postcontroller

import (
	"SociLinkApi/dto"
	frienshiprepository "SociLinkApi/repository/frienship"
	postrepository "SociLinkApi/repository/post"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetOwnTimeline(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	friends, err := frienshiprepository.GetFriendships(userId, db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	userIds := make([]uuid.UUID, len(friends)+1)

	userIds[0] = userId
	for i, friend := range friends {
		id := friend.FriendID

		if friend.FriendID == userId {
			id = friend.UserID
		}

		userIds[i+1] = id
	}

	if posts, err := postrepository.GetPostsByUsers(userIds, db); err != nil {
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
