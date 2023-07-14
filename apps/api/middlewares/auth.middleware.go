package middlewares

import (
	"aurora/services/cache"
	"aurora/services/db"
	"aurora/services/db/models"
	"aurora/services/jwt"
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("x-access-token")

		if accessToken == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Missing access token")
			c.Abort()
			return
		}

		claims, err := jwt.DecodeJwt(accessToken)

		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		cacheResult, err := cache.HGet[jwt.Payload](cache.GetFormattedKey(cache.AuthKeyFormat, claims.Email))

		if err == nil && cacheResult != nil {
			c.Set("user", cacheResult)
			c.Next()
			return
		}

		// Cache miss
		var auth models.Auth
		result := db.Client.Find(&auth, "email = ?", claims.Email)

		if result.Error != nil || auth.Email == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
			c.Abort()
			return
		}

		reqUser := jwt.Payload{
			Id:       auth.ID.String(),
			FullName: auth.FullName,
			Email:    auth.Email,
		}

		c.Set("user", reqUser)
		c.Next()

		// Set cache
		_ = cache.HSet(cache.GetFormattedKey(cache.AuthKeyFormat, claims.Email), reqUser, cache.AuthTTL)
	}
}
