package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/handlers/dto"
	"aurora/services/cache"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func CreateCategory(c *gin.Context) {
	body := c.MustGet("body").(dto.CreateCategoryDto)

	if body.ParentId != nil {
		if _, err := uuid.Parse(body.ParentId.String()); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "parentId is malformed")
			return
		}
	}

	category := &models.Category{
		Name:     body.Name,
		ParentId: body.ParentId,
	}

	res := db.Client.Create(&category)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.Status(http.StatusCreated)
}

func GetCategories(c *gin.Context) {
	cacheRes, err := cache.HGet[dto.MegaNavigation]("mega-navigation")

	if err == nil && cacheRes != nil {
		c.JSON(http.StatusOK, gin.H{
			"data": cacheRes,
		})
		return
	}

	// Cache miss
	var categories []models.Category
	res := db.Client.Find(&categories)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	megaNavigation := getMegaNavigationFromCategories(categories)

	_ = cache.HSet("mega-navigation", megaNavigation, 1*time.Hour)

	c.JSON(http.StatusOK, gin.H{
		"data": megaNavigation,
	})
}

func getMegaNavigationFromCategories(categories []models.Category) dto.MegaNavigation {
	var nav dto.MegaNavigation

	l1Categories := make([]models.Category, 0)
	l2Categories := make([]models.Category, 0)
	uncategorized := make([]models.Category, 0)

	// Find L0 categories
	for _, category := range categories {
		if category.ParentId == nil {
			nav.Items = append(nav.Items, dto.L0Item{
				Id:    category.Id.String(),
				Name:  category.Name,
				Items: make([]dto.L1Item, 0),
			})
		} else {
			uncategorized = append(uncategorized, category)
		}
	}

	// Find L1 and L2 categories
	for _, category := range uncategorized {
		for i, l0Category := range nav.Items {
			if category.ParentId.String() == l0Category.Id {
				l1Categories = append(l1Categories, category)
				nav.Items[i].Items = append(nav.Items[i].Items, dto.L1Item{
					Id:    category.Id.String(),
					Name:  category.Name,
					Items: make([]dto.L2Item, 0),
				})
			} else {
				l2Categories = append(l2Categories, category)
			}
		}
	}

	// Build L2 categories
	for _, category := range l2Categories {
		for i, l0 := range nav.Items {
			for j, l1 := range l0.Items {
				if l1.Id == category.ParentId.String() {
					nav.Items[i].Items[j].Items = append(nav.Items[i].Items[j].Items, dto.L2Item{
						Id:   category.Id.String(),
						Name: category.Name,
					})
				}
			}
		}
	}

	return nav
}
