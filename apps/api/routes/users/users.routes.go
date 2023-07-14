package users

import (
	"aurora/services/aws/models"
	"aurora/services/cache"
	"aurora/services/jwt"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMe(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)
	email := reqUser.Email

	if hit := utils.CheckCache[models.User](c, cache.UserKey(email)); hit {
		return
	}

	// Cache miss
	var user *models.User
	res, err := user.GetUserByEmail(email)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Empty field means user not found
	if res.Id == "" {
		utils.ErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	// Set cache
	utils.SetCache(c, cache.UserKey(email), res, cache.UserTTL)

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
