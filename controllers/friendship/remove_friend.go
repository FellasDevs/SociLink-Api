package friendshipcontroller

import (
	frienshiprepository "SociLinkApi/repository/frienship"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func RemoveFriend(context *gin.Context, db *gorm.DB) {
	uid, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Erro ao obter id do usuário",
		})
		return
	}

	friendshipId, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id do amigo não informado",
		})
		return
	}

	friendship, err := frienshiprepository.GetFriendshipRequestById(friendshipId, db)
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

	if uid.(uuid.UUID) != friendship.UserID && uid.(uuid.UUID) != friendship.FriendID {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Você não tem permissão para remover esta amizade",
		})
		return
	}

	if err := frienshiprepository.DeleteFriendship(friendshipId, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "amizade desfeita com sucesso",
		})
	}
}
