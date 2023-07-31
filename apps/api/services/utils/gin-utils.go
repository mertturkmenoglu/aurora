package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ErrorResponse(c *gin.Context, status int, errorMessage string) {
	c.JSON(status, gin.H{
		"error": errorMessage,
	})
}

func isNotFoundError(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func HandleDatabaseError(c *gin.Context, err error) {
	if isNotFoundError(err) {
		ErrorResponse(c, http.StatusNotFound, "resource not found")
		c.Abort()
		return
	}

	ErrorResponse(c, http.StatusInternalServerError, err.Error())
	c.Abort()
}
