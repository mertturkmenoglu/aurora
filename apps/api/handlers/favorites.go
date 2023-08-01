package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/handlers/dto"
	"aurora/services/jwt"
	"aurora/services/utils"
	"aurora/services/utils/pagination"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"net/http"
)

func GetMyFavorites(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)
	params, err := pagination.GetParamsFromContext(c)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var favorites []*models.Favorite
	var count int64

	res := db.Client.
		Preload(clause.Associations).
		Where("user_id = ?", reqUser.UserId).
		Limit(params.PageSize).
		Offset(params.Offset).
		Find(&favorites).
		Count(&count)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       favorites,
		"pagination": pagination.GetPagination(params, count),
	})
}

func AddFavorite(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)
	body := c.MustGet("body").(dto.AddFavoriteDto)

	productIdAsUUID, err := uuid.Parse(body.ProductId)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "productId is missing or malformed")
		return
	}

	userIdAsUUID, err := uuid.Parse(reqUser.UserId)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "userId is missing or malformed")
		return
	}

	favorite := models.Favorite{
		UserId:    userIdAsUUID,
		ProductId: productIdAsUUID,
	}

	res := db.Client.
		Preload(clause.Associations).
		Create(&favorite)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": favorite,
	})
}

func DeleteFavorite(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)
	id := c.Param("id")

	var favorite *models.Favorite

	res := db.Client.
		Preload(clause.Associations).
		Where("user_id = ? AND id = ?", reqUser.UserId, id).
		Delete(&favorite)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": favorite,
	})
}

func DeleteAllFavorites(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)

	var favorites []*models.Favorite

	res := db.Client.
		Preload(clause.Associations).
		Where("user_id = ?", reqUser.UserId).
		Delete(&favorites)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": favorites,
	})
}
