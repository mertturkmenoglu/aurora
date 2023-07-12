package products

import (
	"aurora/services/aws/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	var product *models.Product
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

	c.JSON(http.StatusOK, gin.H{
		"data": productResult,
	})
}
