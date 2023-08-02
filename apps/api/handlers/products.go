package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/handlers/dto"
	"aurora/services/cache"
	"aurora/services/utils"
	"aurora/services/utils/pagination"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

func checkOnlyOneDefaultVariant(vars []dto.ProductVariantDto) bool {
	counter := 0

	for _, variant := range vars {
		if variant.IsDefault {
			counter++
		}
	}

	if counter == 1 {
		return true
	}

	return false
}

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

	variantDefaultsOk := checkOnlyOneDefaultVariant(body.ProductVariants)

	if !variantDefaultsOk {
		utils.ErrorResponse(c, http.StatusBadRequest, "Only one variant can be default")
		return
	}

	var productId uuid.UUID

	err = db.Client.Transaction(func(tx *gorm.DB) error {
		product := &models.Product{
			Name:        body.Name,
			Description: body.Description,
			IsFeatured:  body.IsFeatured,
			IsNew:       body.IsNew,
			IsOnSale:    body.IsOnSale,
			IsPopular:   body.IsPopular,
			BrandId:     brandIdAsUUID,
			CategoryId:  categoryIdAsUUID,
		}

		res := tx.Create(product)

		if res.Error != nil {
			return res.Error
		}

		productId = product.Id

		for _, variant := range body.ProductVariants {
			styleUUID, err := uuid.Parse(variant.ProductStyleId)

			if err != nil {
				return err
			}

			sizeUUID, err := uuid.Parse(variant.ProductSizeId)

			if err != nil {
				return err
			}

			productVariant := &models.ProductVariant{
				ProductId:    product.Id,
				CurrentPrice: variant.CurrentPrice,
				OldPrice:     variant.OldPrice,
				Inventory:    variant.Inventory,
				Image: models.ProductImage{
					Url:       variant.Image.Url,
					ProductId: product.Id,
				},
				ShippingPrice:  variant.ShippingPrice,
				ShippingTime:   variant.ShippingTime,
				ShippingType:   variant.ShippingType,
				ProductStyleId: styleUUID,
				ProductSizeId:  sizeUUID,
			}

			res = tx.Create(productVariant)

			if res.Error != nil {
				return res.Error
			}

			productVariant.ImageId = productVariant.Image.Id

			res = tx.Save(productVariant)

			if res.Error != nil {
				return res.Error
			}

			if variant.IsDefault {
				product.DefaultVariantId = productVariant.Id

				res = tx.Save(product)

				if res.Error != nil {
					return res.Error
				}
			}
		}

		return nil
	})

	if err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	var product *models.Product

	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		First(&product, "ID = ?", productId)

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
		Preload("DefaultVariant.Image").
		Preload("DefaultVariant.ProductStyle").
		Preload("DefaultVariant.ProductSize").
		Preload("ProductVariants.Image").
		Preload("ProductVariants.ProductStyle").
		Preload("ProductVariants.ProductSize").
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

func GetProductsByCategory(c *gin.Context) {
	categoryId := c.Query("categoryId")
	params, err := pagination.GetParamsFromContext(c)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := uuid.Parse(categoryId); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "categoryId is malformed")
		return
	}

	categoryIds := getCategoryAndSubCategoryIds(categoryId)

	var products []*models.Product
	var count int64

	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Preload("DefaultVariant.Image").
		Preload("DefaultVariant.ProductStyle").
		Preload("DefaultVariant.ProductSize").
		Limit(params.PageSize).
		Offset(params.Offset).
		Find(&products, "category_id IN ?", categoryIds).
		Count(&count)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       products,
		"pagination": pagination.GetPagination(params, count),
	})
}

func GetFeaturedProducts(c *gin.Context) {
	var products []*models.Product
	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Preload("DefaultVariant.Image").
		Preload("DefaultVariant.ProductStyle").
		Preload("DefaultVariant.ProductSize").
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
		Preload("DefaultVariant.Image").
		Preload("DefaultVariant.ProductStyle").
		Preload("DefaultVariant.ProductSize").
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
		Preload("DefaultVariant.Image").
		Preload("DefaultVariant.ProductStyle").
		Preload("DefaultVariant.ProductSize").
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
		Preload("DefaultVariant.Image").
		Preload("DefaultVariant.ProductStyle").
		Preload("DefaultVariant.ProductSize").
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

func GetAllProducts(c *gin.Context) {
	var products []*models.Product
	params, err := pagination.GetParamsFromContext(c)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var count int64

	res := db.Client.
		Preload(clause.Associations).
		Preload("Category.Parent").
		Preload("Category.Parent.Parent").
		Preload("DefaultVariant.Image").
		Preload("DefaultVariant.ProductStyle").
		Preload("DefaultVariant.ProductSize").
		Limit(params.PageSize).
		Offset(params.Offset).
		Find(&products).
		Count(&count)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       products,
		"pagination": pagination.GetPagination(params, count),
	})
}

func AddProductStyles(c *gin.Context) {
	body := c.MustGet("body").(dto.AddProductStylesDto)
	styles := make([]*models.ProductStyle, len(body.Styles))

	for i, style := range body.Styles {
		styles[i] = &models.ProductStyle{
			Name: style.Name,
		}
	}

	res := db.Client.Create(&styles)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": styles,
	})
}

func AddProductSizes(c *gin.Context) {
	body := c.MustGet("body").(dto.AddProductSizesDto)
	sizes := make([]*models.ProductSize, len(body.Sizes))

	for i, size := range body.Sizes {
		sizes[i] = &models.ProductSize{
			Name: size.Name,
		}
	}

	res := db.Client.Create(&sizes)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": sizes,
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
