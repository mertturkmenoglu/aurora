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
