package middlewares

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/services/jwt"
	"aurora/services/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminKey, ok := os.LookupEnv("ADMIN_KEY")

		if !ok {
			panic("ADMIN_KEY not found")
		}

		accessToken := c.GetHeader("x-access-token")
		adminKeyFromHeader := c.GetHeader("x-admin-key")

		if accessToken == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Not authorized")
			c.Abort()
			return
		}

		if adminKeyFromHeader != adminKey {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Not authorized")
			c.Abort()
			return
		}

		claims, err := jwt.DecodeJwt(accessToken)

		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		var user models.User
		result := db.Client.First(&user, "email = ?", claims.Email)

		if result.Error != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Not authorized")
			c.Abort()
			return
		}

		var admin models.Admin
		result = db.Client.First(&admin, "user_id = ?", user.Id)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				utils.ErrorResponse(c, http.StatusUnauthorized, "Not authorized")
				c.Abort()
				return
			}

			utils.ErrorResponse(c, http.StatusInternalServerError, result.Error.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}
