package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/services/utils"
	"aurora/services/utils/pagination"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"net/http"
)

func SearchProducts(c *gin.Context) {
	searchTerm, ok := c.GetQuery("q")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Search term is required",
		})
	}

	params, err := pagination.GetParamsFromContext(c)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var products []*models.Product
	var count int64

	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Limit(params.PageSize).
		Offset(params.Offset).
		Where("name ILIKE ?", "%"+searchTerm+"%").
		Find(&products).
		Count(&count)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(200, gin.H{
		"data":       products,
		"pagination": pagination.GetPagination(params, count),
	})
}
