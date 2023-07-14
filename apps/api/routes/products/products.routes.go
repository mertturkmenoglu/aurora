package products

import (
	"aurora/services/db"
	"aurora/services/db/models"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	var product *models.Product
	result := db.Client.First(&product, "ID = ?", id)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "record not found") {
			utils.ErrorResponse(c, http.StatusNotFound, "product not found")
			return
		}

		utils.ErrorResponse(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}
