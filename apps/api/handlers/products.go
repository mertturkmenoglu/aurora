package handlers

import (
	"aurora/services/cache"
	"aurora/services/db"
	"aurora/services/db/models"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	key := cache.GetFormattedKey(cache.ProductKeyFormat, id)
	cacheResult, err := cache.HGet[models.Product](key)

	if cacheResult != nil && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": cacheResult,
		})
		return
	}

	var product *models.Product
	res := db.Client.First(&product, "ID = ?", id)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	_ = cache.HSet(key, product, cache.ProductTTL)

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}
