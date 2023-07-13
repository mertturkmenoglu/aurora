package utils

import (
	"aurora/services/aws/models"
	"aurora/services/cache"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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

func SetCache[T models.DynamoModel](c *gin.Context, key string, data *T, ttl time.Duration) {
	serializedData, err := json.Marshal(data)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	obj := map[string]string{
		"data": string(serializedData),
	}

	err = cache.HSet(key, obj, ttl)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
}