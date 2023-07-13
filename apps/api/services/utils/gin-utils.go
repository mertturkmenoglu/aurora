package utils

import (
	"github.com/gin-gonic/gin"
)

func ErrorResponse(c *gin.Context, status int, errorMessage string) {
	c.JSON(status, gin.H{
		"error": errorMessage,
	})
}
