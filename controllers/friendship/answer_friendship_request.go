package friendshipcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	friendshiprepository "SociLinkApi/repository/friendship"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func AnswerFriendshipRequest(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	var params dto.AnswerFriendshipRequestDto
	if err := context.ShouldBindJSON(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	requestId, err := uuid.Parse(params.RequestId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	friendship := models.Friendship{ID: requestId}
	if err = friendshiprepository.GetFriendship(&friendship, db); err != nil {
		var statusCode int

		if err.Error() == "record not found" {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}

		context.JSON(statusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	if friendship.FriendID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Você não tem permissão para responder a este pedido de amizade",
		})
		return
	}

	if !friendship.Pending && friendship.Accepted {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Este pedido de amizade já foi aceito",
		})
		return
	}

	friendship.Accepted = params.Answer
	friendship.Pending = false

	if err := friendshiprepository.UpdateFriendshipRequest(friendship, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "pedido de amizade respondido com sucesso",
		})
	}
}
