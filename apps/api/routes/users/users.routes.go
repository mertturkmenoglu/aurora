package users

import (
	"aurora/services/aws/models"
	"aurora/services/cache"
	"aurora/services/jwt"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserById(c *gin.Context) {
	email := c.Param("email")
	reqUser := c.MustGet("user").(jwt.Payload)

	// Check if the user is requesting their own data
	if email != reqUser.Email {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	utils.CheckCache[models.User](c, cache.UserKey(email))

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
	utils.SetCache(c, cache.UserKey(email), res)

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
