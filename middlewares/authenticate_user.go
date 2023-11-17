package middlewares

import (
	"SociLinkApi/dto"
	authservice "SociLinkApi/services/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

func AuthenticateUser(context *gin.Context) {
	header := dto.AuthHeader{}

	if err := context.ShouldBindHeader(&header); err != nil || header.AuthToken == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Token de autorização não informado.",
		})
		return
	}

	token := strings.Split(header.AuthToken, "Bearer ")

	if len(token) != 2 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Token de autorização inválido.",
		})
		return
	}

	claims, err := authservice.ParseAuthToken(token[1])
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	userId, err := uuid.Parse(claims.UserId)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	context.Set("userId", userId)
}
