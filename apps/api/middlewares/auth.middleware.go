package middlewares

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/services/cache"
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

		key := cache.GetFormattedKey(cache.AuthKeyFormat, claims.Email)
		cacheResult, err := cache.HGet[jwt.Payload](key)

		if err == nil && cacheResult != nil {
			c.Set("user", *cacheResult)
			c.Next()
			return
		}

		// Cache miss
		var user models.User
		result := db.Client.Find(&user, "email = ?", claims.Email)

		if result.Error != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
			c.Abort()
			return
		}

		reqUser := jwt.Payload{
			UserId:   user.Id.String(),
			FullName: user.FullName,
			Email:    user.Email,
		}

		// Set cache
		_ = cache.HSet(key, reqUser, cache.AuthTTL)

		c.Set("user", reqUser)
		c.Next()
	}
}
