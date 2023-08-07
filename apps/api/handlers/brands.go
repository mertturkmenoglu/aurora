package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/handlers/dto"
	"aurora/services/cache"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateBrand(c *gin.Context) {
	body := c.MustGet("body").(dto.CreateBrandDto)
	brand := models.Brand{
		Name:        body.Name,
		Description: body.Description,
	}

	res := db.Client.Create(&brand)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": brand,
	})
}

func GetBrandById(c *gin.Context) {
	id, err := utils.GetIdFromParams(c)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	key := cache.GetFormattedKey(cache.BrandKeyFormat, id.String())
	cacheResult, err := cache.HGet[models.Brand](key)

	if cacheResult != nil && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": cacheResult,
		})
		return
	}

	var brand models.Brand

	res := db.Client.First(&brand, id)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	_ = cache.HSet(key, brand, cache.BrandTTL)

	c.JSON(http.StatusOK, gin.H{
		"data": brand,
	})
}
