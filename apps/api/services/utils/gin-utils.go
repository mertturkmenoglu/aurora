package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ErrorResponse(c *gin.Context, status int, errorMessage string) {
	c.JSON(status, gin.H{
		"error": errorMessage,
	})
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "record not found")
}

func HandleDatabaseError(c *gin.Context, err error) {
	if isNotFoundError(err) {
		ErrorResponse(c, http.StatusNotFound, "record not found")
		c.Abort()
		return
	}

	ErrorResponse(c, http.StatusInternalServerError, err.Error())
	c.Abort()
}
