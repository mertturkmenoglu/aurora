package products

import (
	"aurora/services/aws/models"
	"aurora/services/cache"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	var product *models.Product

	if hit := utils.CheckCache[models.Product](c, cache.ProductKey(id)); hit {
		return
	}

	// Cache miss
	productResult, err := product.GetProductById(id)

	if err != nil {
		if productResult.Id == "" {
			utils.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Check if product exists
	if productResult.Id == "" {
		utils.ErrorResponse(c, http.StatusNotFound, "product not found")
		return
	}

	// Set cache
	utils.SetCache[models.Product](c, cache.ProductKey(id), productResult, cache.ProductTTL)

	c.JSON(http.StatusOK, gin.H{
		"data": productResult,
	})
}
