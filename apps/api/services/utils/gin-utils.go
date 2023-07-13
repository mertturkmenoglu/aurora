package utils

import (
	"aurora/services/aws/models"
	"aurora/services/cache"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorResponse(c *gin.Context, status int, errorMessage string) {
	c.JSON(status, gin.H{
		"error": errorMessage,
	})
}

func CheckCache[T models.DynamoModel](c *gin.Context, key string) {
	var result *T

	cacheResult, err := cache.HGetAll(key)

	// Cache hit
	if err == nil && len(cacheResult) > 0 {
		err = json.Unmarshal([]byte(cacheResult["data"]), &result)

		if err != nil {
			ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
		return
	}

	// Cache miss, continue to regular flow
}
