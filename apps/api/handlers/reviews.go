package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/handlers/dto"
	"aurora/services/jwt"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func CreateBrandReview(c *gin.Context) {
	body := c.MustGet("body").(dto.CreateBrandReviewDto)
	reqUser := c.MustGet("user").(jwt.Payload)

	brandIdAsUuid, err := uuid.Parse(body.BrandId)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "brandId is missing or malformed")
		return
	}

	var user *models.User
	res := db.Client.
		First(&user, "email = ?", reqUser.Email)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	brandReview := models.BrandReview{
		Comment: body.Comment,
		Rating:  body.Rating,
		BrandId: brandIdAsUuid,
		UserId:  user.Id,
	}

	res = db.Client.Create(&brandReview)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": brandReview,
	})
}

func CreateProductReview(c *gin.Context) {
	body := c.MustGet("body").(dto.CreateProductReviewDto)
	reqUser := c.MustGet("user").(jwt.Payload)

	_, err := uuid.Parse(body.ProductId)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "productId is missing or malformed")
		return
	}

	var user *models.User

	res := db.Client.
		First(&user, "email = ?", reqUser.Email)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}
}

func GetMyBrandReviews(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)

	var brandReviews []*models.BrandReview

	res := db.Client.
		Preload("Brand").
		Where("user_id = ?", reqUser.Id).
		Find(&brandReviews)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": brandReviews,
	})
}

func GetMyProductReviews(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)

	var productReviews []*models.ProductReview

	res := db.Client.
		Preload("Product").
		Where("user_id = ?", reqUser.Id).
		Find(&productReviews)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productReviews,
	})
}

func GetBrandReview(c *gin.Context) {
	id := c.Param("id")

	idAsUUID, err := uuid.Parse(id)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is missing or malformed")
		return
	}

	var brandReview *models.BrandReview

	res := db.Client.
		Preload("Brand").
		First(&brandReview, "id = ?", idAsUUID)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": brandReview,
	})
}

func GetProductReview(c *gin.Context) {
	id := c.Param("id")

	idAsUUID, err := uuid.Parse(id)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is missing or malformed")
		return
	}

	var productReview *models.ProductReview

	res := db.Client.
		Preload("Product").
		First(&productReview, "id = ?", idAsUUID)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productReview,
	})
}

func GetBrandReviews(c *gin.Context) {
	id := c.Param("id")

	idAsUUID, err := uuid.Parse(id)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is missing or malformed")
		return
	}

	var brandReviews []*models.BrandReview

	res := db.Client.
		Preload("User").
		Where("brand_id = ?", idAsUUID).
		Find(&brandReviews)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": brandReviews,
	})
}

func GetProductReviews(c *gin.Context) {
	id := c.Param("id")

	idAsUUID, err := uuid.Parse(id)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is missing or malformed")
		return
	}

	var productReviews []*models.ProductReview

	res := db.Client.
		Preload("User").
		Where("product_id = ?", idAsUUID).
		Find(&productReviews)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": productReviews,
	})
}

func DeleteBrandReview(c *gin.Context) {
	id := c.Param("id")

	idAsUUID, err := uuid.Parse(id)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is missing or malformed")
		return
	}

	reqUser := c.MustGet("user").(jwt.Payload)
	userIdAsUUID, err := uuid.Parse(reqUser.Id)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is missing or malformed")
		return
	}

	var brandReview *models.BrandReview

	res := db.Client.
		First(&brandReview, "id = ?", idAsUUID)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	if brandReview.UserId != userIdAsUUID {
		utils.ErrorResponse(c, http.StatusForbidden, "you are not the owner of this review")
		return
	}

	res = db.Client.Delete(&brandReview)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "review deleted successfully",
	})
}

func DeleteProductReview(c *gin.Context) {
	id := c.Param("id")

	idAsUUID, err := uuid.Parse(id)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is missing or malformed")
		return
	}

	reqUser := c.MustGet("user").(jwt.Payload)
	userIdAsUUID, err := uuid.Parse(reqUser.Id)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "id is missing or malformed")
		return
	}

	var productReview *models.ProductReview

	res := db.Client.
		First(&productReview, "id = ?", idAsUUID)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	if productReview.UserId != userIdAsUUID {
		utils.ErrorResponse(c, http.StatusForbidden, "you are not the owner of this review")
		return
	}

	res = db.Client.Delete(&productReview)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "review deleted successfully",
	})
}
