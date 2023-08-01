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

	var user *models.User

	res := db.Client.
		First(&user, "email = ?", reqUser.Email)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	var favorites []*models.Favorite
	var count int64

	res = db.Client.
		Preload(clause.Associations).
		Where("user_id = ?", user.Id).
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

	var user *models.User

	res := db.Client.
		First(&user, "email = ?", reqUser.Email)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	favorite := models.Favorite{
		UserId:    user.Id,
		ProductId: productIdAsUUID,
	}

	res = db.Client.
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

	var user *models.User

	res := db.Client.
		First(&user, "email = ?", reqUser.Email)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	var favorite *models.Favorite

	res = db.Client.
		Preload(clause.Associations).
		Where("user_id = ? AND id = ?", user.Id, id).
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

	var user *models.User

	res := db.Client.
		First(&user, "email = ?", reqUser.Email)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	var favorites []*models.Favorite

	res = db.Client.
		Preload(clause.Associations).
		Where("user_id = ?", user.Id).
		Delete(&favorites)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": favorites,
	})
}
