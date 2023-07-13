package brands

import (
	"aurora/services/aws/models"
	"aurora/services/cache"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBrandById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is required")
		return
	}

	var brand models.Brand

	if hit := utils.CheckCache[models.Brand](c, cache.BrandKey(id)); hit {
		return
	}

	// Cache miss
	brandResult, err := brand.GetBrandById(id)

	if err != nil {
		if brandResult.Id == "" {
			utils.ErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Check if brand exists
	if brandResult.Id == "" {
		utils.ErrorResponse(c, http.StatusNotFound, "brand not found")
		return
	}

	// Set cache
	utils.SetCache[models.Brand](c, cache.BrandKey(id), brandResult, cache.BrandTTL)

	c.JSON(http.StatusOK, gin.H{
		"data": brandResult,
	})
}
