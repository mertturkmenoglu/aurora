package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func GetIdFromParams(c *gin.Context) (uuid.UUID, error) {
	id, exists := c.Params.Get("id")

	if !exists {
		return uuid.New(), errors.New("id is missing")
	}

	u, err := uuid.Parse(id)

	if err != nil {
		return uuid.New(), errors.New("id is malformed")
	}

	return u, nil
}
