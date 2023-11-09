package friendshipcontroller

import (
	"SociLinkApi/dto"
	frienshiprepository "SociLinkApi/repository/frienship"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func AnswerFriendshipRequest(context *gin.Context, db *gorm.DB) {
	uid, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Erro ao obter id do usuário",
		})
		return
	}

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

	friendship, err := frienshiprepository.GetFriendshipRequestById(requestId, db)
	if err != nil {
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

	if friendship.FriendID != uid.(uuid.UUID) {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Você não tem permissão para responder a este pedido de amizade",
		})
		return
	}

	if friendship.Pending == false {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Este pedido de amizade já foi respondido",
		})
		return
	}

	friendship.Accepted = params.Answer
	friendship.Pending = false

	if err := frienshiprepository.UpdateFriendshipRequest(friendship, db); err != nil {
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
