package friendshipcontroller

import (
	userrepository "SociLinkApi/repository/frienship"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func RequestFriendship(context *gin.Context, db *gorm.DB) {
	uid, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Erro ao obter id do usuário",
		})
		return
	}

	friendId, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id do amigo inválido",
		})
		return
	}

	if uid.(uuid.UUID) == friendId {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Você não pode enviar uma solicitação de amizade a si mesmo",
		})
		return
	}

	if err := userrepository.CreateFriendshipRequest(uid.(uuid.UUID), friendId, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "pedido de amizade enviado com sucesso",
		})
	}
}
