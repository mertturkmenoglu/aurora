package users

import (
	"aurora/services/aws/models"
	"aurora/services/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserById(c *gin.Context) {
	email := c.Param("email")
	reqUser := c.MustGet("user").(jwt.Payload)

	// Check if the user is requesting their own data
	if email != reqUser.Email {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	var user *models.User

	res, err := user.GetUserByEmail(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Empty field means user not found
	if res.Id == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
