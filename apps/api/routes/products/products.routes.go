package products

import (
	"aurora/services/aws/models"
	"aurora/services/cache"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	var product *models.Product

	// Check cache
	cacheResult, err := cache.HGetAll(cache.ProductKey(id))

	// Cache hit
	if err == nil && len(cacheResult) > 0 {
		err = json.Unmarshal([]byte(cacheResult["data"]), &product)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": product,
		})
		return
	}

	// Cache miss
	productResult, err := product.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Check if product exists
	if productResult.Id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Product not found",
		})
		return
	}

	// Set cache
	serializedProduct, err := json.Marshal(productResult)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = cache.HSet(cache.ProductKey(id), map[string]string{
		"data": string(serializedProduct),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productResult,
	})
}
