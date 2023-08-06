package handlers

import (
	"aurora/db/queries"
	"aurora/services/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHomeAggregation(c *gin.Context) {
	featuredProducts, featuredProductsErr := queries.GetFeaturedProducts()
	newProducts, newProductsErr := queries.GetNewProducts()
	saleProducts, saleProductsErr := queries.GetSaleProducts()
	popularProducts, popularProductsErr := queries.GetPopularProducts()

	err := errors.Join(featuredProductsErr, newProductsErr, saleProductsErr, popularProductsErr)

	if err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"featured": featuredProducts,
			"new":      newProducts,
			"sale":     saleProducts,
			"popular":  popularProducts,
		},
	})
}
