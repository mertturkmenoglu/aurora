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
	var megaNavigation dto.MegaNavigation

	res := db.Client.Find(&categories)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	l1Categories := make([]models.Category, 0)
	l2Categories := make([]models.Category, 0)

	for _, category := range categories {
		if category.ParentId == nil {
			megaNavigation.Items = append(megaNavigation.Items, dto.L0Item{
				Id:    category.Id.String(),
				Name:  category.Name,
				Items: make([]dto.L1Item, 0),
			})
		}
	}

	for _, category := range categories {
		for i, l0Category := range megaNavigation.Items {
			if category.ParentId != nil && category.ParentId.String() == l0Category.Id {
				l1Categories = append(l1Categories, category)
				megaNavigation.Items[i].Items = append(megaNavigation.Items[i].Items, dto.L1Item{
					Id:    category.Id.String(),
					Name:  category.Name,
					Items: make([]dto.L2Item, 0),
				})
			}
		}
	}

	for _, category := range categories {
		for _, l1Category := range l1Categories {
			if category.ParentId != nil && category.ParentId.String() == l1Category.Id.String() {
				l2Categories = append(l2Categories, category)
			}
		}
	}

	for _, l2 := range l2Categories {
		for i, l0 := range megaNavigation.Items {
			for j, l1 := range l0.Items {
				if l1.Id == l2.ParentId.String() {
					megaNavigation.Items[i].Items[j].Items = append(megaNavigation.Items[i].Items[j].Items, dto.L2Item{
						Id:   l2.Id.String(),
						Name: l2.Name,
					})
				}
			}
		}
	}

	_ = cache.HSet("mega-navigation", megaNavigation, 1*time.Hour)

	c.JSON(http.StatusOK, gin.H{
		"data": megaNavigation,
	})
}
