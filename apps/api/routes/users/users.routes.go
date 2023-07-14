package users

import (
	_ "aurora/services/cache"
	"aurora/services/db"
	"aurora/services/db/models"
	"aurora/services/jwt"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetMe(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)
	email := reqUser.Email

	var user *models.User
	result := db.Client.Find(&user, "email = ?", email)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "record not found") {
			utils.ErrorResponse(c, http.StatusNotFound, "User not found")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
