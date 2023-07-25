package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/handlers/dto"
	"aurora/services/cache"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"net/http"
)

func CreateProduct(c *gin.Context) {
	body := c.MustGet("body").(dto.CreateProductDto)

	brandIdAsUUID, err := uuid.Parse(body.BrandId)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "brandId is missing or malformed")
		return
	}

	categoryIdAsUUID, err := uuid.Parse(body.CategoryId)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "categoryId is missing or malformed")
		return
	}

	product := &models.Product{
		Name:          body.Name,
		Description:   body.Description,
		CurrentPrice:  body.CurrentPrice,
		OldPrice:      body.OldPrice,
		Inventory:     body.Inventory,
		IsFeatured:    body.IsFeatured,
		IsNew:         body.IsNew,
		IsOnSale:      body.IsOnSale,
		IsPopular:     body.IsPopular,
		ShippingPrice: body.ShippingPrice,
		ShippingTime:  body.ShippingTime,
		ShippingType:  body.ShippingType,
		Slug:          body.Slug,
		BrandId:       brandIdAsUUID,
		CategoryId:    categoryIdAsUUID,
		Images:        []models.ProductImage{},
		ProductStyles: []models.ProductStyle{},
		ProductSizes:  []models.ProductSize{},
	}

	res := db.Client.Create(product)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	images := make([]*models.ProductImage, len(body.Images))

	for i, image := range body.Images {
		images[i] = &models.ProductImage{
			ProductId: product.Id,
			Url:       image.Url,
		}
	}

	res = db.Client.Create(&images)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	styles := make([]*models.ProductStyle, len(body.ProductStyles))

	for i, style := range body.ProductStyles {
		styles[i] = &models.ProductStyle{
			ProductId: product.Id,
			Name:      style.Name,
		}
	}

	res = db.Client.Create(&styles)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	sizes := make([]*models.ProductSize, len(body.ProductSizes))

	for i, size := range body.ProductSizes {
		sizes[i] = &models.ProductSize{
			ProductId: product.Id,
			Name:      size.Name,
		}
	}

	res = db.Client.Create(&sizes)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": product,
	})
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is malformed")
		return
	}

	key := cache.GetFormattedKey(cache.ProductKeyFormat, id)
	cacheResult, err := cache.HGet[models.Product](key)

	if cacheResult != nil && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": cacheResult,
		})
		return
	}

	var product *models.Product
	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		First(&product, "ID = ?", id)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	_ = cache.HSet(key, product, cache.ProductTTL)

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func GetProductByCategory(c *gin.Context) {
	categoryId := c.Query("categoryId")

	if _, err := uuid.Parse(categoryId); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "categoryId is malformed")
		return
	}

	categoryIds := getCategoryAndSubCategoryIds(categoryId)

	var products []*models.Product
	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Find(&products, "category_id IN ?", categoryIds)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func GetFeaturedProducts(c *gin.Context) {
	var products []*models.Product
	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Order("created_at desc").
		Limit(25).
		Find(&products, "is_featured = ?", true)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func GetNewProducts(c *gin.Context) {
	var products []*models.Product
	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Order("created_at desc").
		Limit(25).
		Find(&products, "is_new = ?", true)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func GetSaleProducts(c *gin.Context) {
	var products []*models.Product
	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Order("created_at desc").
		Limit(25).
		Find(&products, "is_on_sale = ?", true)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func GetPopularProducts(c *gin.Context) {
	var products []*models.Product
	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Order("created_at desc").
		Limit(25).
		Find(&products, "is_popular = ?", true)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func GetFreeShippingProducts(c *gin.Context) {
	var products []*models.Product
	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Order("created_at desc").
		Limit(25).
		Find(&products, "shipping_price = ?", 0)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func GetAllProducts(c *gin.Context) {
	var products []*models.Product

	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Find(&products)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func getCategoryAndSubCategoryIds(categoryId string) []string {
	var categories []*models.Category
	categoryIds := []string{categoryId}

	for i := 0; i < 3; i++ {
		var tmp []*models.Category

		db.Client.Find(&tmp, "parent_id IN ?", categoryIds)
		categories = append(categories, tmp...)
		ids := make([]string, len(tmp))

		for j, category := range tmp {
			ids[j] = category.Id.String()
		}

		categoryIds = ids
	}

	categoryIds = make([]string, len(categories))

	for i, category := range categories {
		categoryIds[i] = category.Id.String()
	}

	categoryIds = append(categoryIds, categoryId)

	return categoryIds
}
