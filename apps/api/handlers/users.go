package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/handlers/dto"
	"aurora/services/cache"
	"aurora/services/jwt"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMe(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)
	email := reqUser.Email

	key := cache.GetFormattedKey(cache.UserKeyFormat, email)
	cacheResult, err := cache.HGet[models.User](key)

	if cacheResult != nil && err == nil {
		c.JSON(http.StatusOK, gin.H{
			"data": cacheResult,
		})
		return
	}

	var user *models.User
	res := db.Client.
		Preload("AdPreference").
		Preload("Addresses").
		Find(&user, "email = ?", email)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	_ = cache.HSet(key, user, cache.UserTTL)

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func UpdateMe(c *gin.Context) {
	body := c.MustGet("body").(dto.UpdateUserDto)
	reqUser := c.MustGet("user").(jwt.Payload)

	res := db.Client.
		Model(&models.User{}).
		Where("email = ?", reqUser.Email).
		Updates(body)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}

func GetMyAddresses(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)

	var addresses []*models.Address

	res := db.Client.
		Where("user_id = ?", reqUser.UserId).
		Find(&addresses)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": addresses,
	})
}

func AddAddress(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

func UpdateAddress(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

func DeleteAddress(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

func UpdateAdPreferences(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}
