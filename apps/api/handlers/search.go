package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func SearchProducts(c *gin.Context) {
	searchTerm := c.Query("q")

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
