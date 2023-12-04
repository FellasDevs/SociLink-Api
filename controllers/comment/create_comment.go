package commentcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	commentrepository "SociLinkApi/repository/comment"
	postrepository "SociLinkApi/repository/post"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateComment(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	var params dto.CreateCommentRequestDto
	if err := context.ShouldBindJSON(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Parâmetros inválidos",
		})
		return
	}

	postUuid, err := uuid.Parse(params.PostId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Id do post inválido",
		})
		return
	}

	if err = postrepository.GetPost(&models.Post{ID: postUuid}, &userId, db); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Post não encontrado",
		})
		return
	}

	comment := models.Comment{
		UserID:  userId,
		PostID:  postUuid,
		Content: params.Content,
	}

	if err = commentrepository.CreateComment(&comment, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Erro ao adicionar o comentário",
		})
		return
	} else {
		context.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Comentário adicionado com sucesso",
		})
		return
	}
}
