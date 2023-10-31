package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendUnknownError(err error, context *gin.Context) {
	context.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"message": err.Error(),
	})
}
