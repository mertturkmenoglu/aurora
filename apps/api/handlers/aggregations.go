package handlers

import (
	"aurora/db/models"
	"aurora/db/queries"
	"aurora/services/cache"
	"aurora/services/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeAggregation struct {
	FeaturedProducts []*models.Product `json:"featured"`
	NewProducts      []*models.Product `json:"new"`
	SaleProducts     []*models.Product `json:"sale"`
	PopularProducts  []*models.Product `json:"popular"`
}

func GetHomeAggregation(c *gin.Context) {
	key := cache.HomeAggregationKey
	cacheResult, err := cache.HGet[HomeAggregation](key)

	if cacheResult != nil && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": cacheResult,
		})
		return
	}

	// Cache miss
	featuredProducts, featuredProductsErr := queries.GetFeaturedProducts()
	newProducts, newProductsErr := queries.GetNewProducts()
	saleProducts, saleProductsErr := queries.GetSaleProducts()
	popularProducts, popularProductsErr := queries.GetPopularProducts()

	err = errors.Join(featuredProductsErr, newProductsErr, saleProductsErr, popularProductsErr)

	if err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	homeAggregation := HomeAggregation{
		FeaturedProducts: featuredProducts,
		NewProducts:      newProducts,
		SaleProducts:     saleProducts,
		PopularProducts:  popularProducts,
	}

	_ = cache.HSet(key, homeAggregation, cache.HomeAggregationTTL)

	c.JSON(http.StatusOK, gin.H{
		"data": homeAggregation,
	})
}
