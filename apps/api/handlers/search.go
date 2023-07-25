package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

func SearchProducts(c *gin.Context) {
	searchTerm := c.Query("q")

	if searchTerm == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Search term is required",
		})
	}

	var products []*models.Product

	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Where("name ILIKE ?", "%"+searchTerm+"%").
		Find(&products)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(200, gin.H{
		"data": products,
	})
}
