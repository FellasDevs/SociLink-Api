package friendshipcontroller

import (
	userrepository "SociLinkApi/repository/friendship"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func RequestFriendship(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	friendId, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id do amigo inválido",
		})
		return
	}

	if userId == friendId {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Você não pode enviar uma solicitação de amizade a si mesmo",
		})
		return
	}

	if friendship, err := userrepository.GetFriendshipByUsers(userId, friendId, db); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
	} else {
		message := "Essa pessoa já enviou uma solicitação de amizade para você"

		if friendship.Accepted {
			message = "Você já é amigo desta pessoa"
		} else if friendship.UserID == userId {
			message = "Você já enviou uma solicitação de amizade para esta pessoa"
		}

		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": message,
		})
		return
	}

	if err = userrepository.CreateFriendshipRequest(userId, friendId, db); err != nil {
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
