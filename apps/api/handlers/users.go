package handlers

import (
	"aurora/services/cache"
	"aurora/services/db"
	"aurora/services/db/models"
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
