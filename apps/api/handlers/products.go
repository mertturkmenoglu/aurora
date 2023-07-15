package handlers

import (
	"aurora/handlers/dto"
	"aurora/services/cache"
	"aurora/services/db"
	"aurora/services/db/models"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	body := c.MustGet("body").(dto.CreateProductDto)

	brandIdAsUUID, err := uuid.Parse(body.BrandId)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "brandId is missing or malformed")
		return
	}

	product := &models.Product{
		Name:          body.Name,
		Description:   body.Description,
		CurrentPrice:  body.CurrentPrice,
		OldPrice:      body.OldPrice,
		Inventory:     body.Inventory,
		IsFeatured:    body.IsFeatured,
		IsNew:         body.IsNew,
		IsOnSale:      body.IsOnSale,
		IsPopular:     body.IsPopular,
		ShippingPrice: body.ShippingPrice,
		ShippingTime:  body.ShippingTime,
		ShippingType:  body.ShippingType,
		Slug:          body.Slug,
		BrandId:       brandIdAsUUID,
		Images:        []models.ProductImage{},
	}

	res := db.Client.Create(product)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	images := make([]*models.ProductImage, len(body.Images))

	for i, image := range body.Images {
		images[i] = &models.ProductImage{
			ProductId: product.Id,
			Url:       image.Url,
		}
	}

	res = db.Client.Create(&images)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": product,
	})
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is malformed")
		return
	}

	key := cache.GetFormattedKey(cache.ProductKeyFormat, id)
	cacheResult, err := cache.HGet[models.Product](key)

	if cacheResult != nil && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": cacheResult,
		})
		return
	}

	var product *models.Product
	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		First(&product, "ID = ?", id)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	_ = cache.HSet(key, product, cache.ProductTTL)

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}
