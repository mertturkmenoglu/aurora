package handlers

import (
	"aurora/db/queries"
	"aurora/handlers/dto"
	"aurora/services/cache"
	"aurora/services/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHomeAggregation(c *gin.Context) {
	key := cache.HomeAggregationKey
	cacheResult, err := cache.HGet[dto.HomeAggregation](key)

	if cacheResult != nil && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": cacheResult,
		})
		return
	}

	data := dto.HomeAggregation{}

	featured, featuredErr := queries.GetFeaturedProducts()
	newProducts, newErr := queries.GetNewProducts()
	sale, saleErr := queries.GetSaleProducts()
	popular, popularErr := queries.GetPopularProducts()

	err = errors.Join(featuredErr, newErr, saleErr, popularErr)

	if err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	data.FeaturedProducts = featured
	data.NewProducts = newProducts
	data.SaleProducts = sale
	data.PopularProducts = popular

	_ = cache.HSet(key, data, cache.HomeAggregationTTL)

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
