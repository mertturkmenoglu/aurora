package utils

import (
	"aurora/services/aws/models"
	"aurora/services/cache"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CheckCache[T models.DynamoModel](c *gin.Context, key string) bool {
	var result *T

	cacheResult, err := cache.HGetAll(key)

	log.Println("Cache result: ", cacheResult, err)

	// Cache hit
	if err == nil && len(cacheResult) > 0 {
		err = json.Unmarshal([]byte(cacheResult["data"]), &result)

		if err != nil {
			ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return false
		}

		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
		c.Done()
		return true
	}

	// Cache miss, continue to regular flow
	return false
}

func SetCache[T models.DynamoModel](c *gin.Context, key string, data *T, ttl time.Duration) {
	serializedData, err := json.Marshal(data)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	obj := map[string]string{
		"data": string(serializedData),
	}

	err = cache.HSet(key, obj, ttl)

	if err != nil {
		ErrorResponse(c, http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}
}
