package middlewares

import (
	"aurora/services/cache"
	"aurora/services/db"
	"aurora/services/db/models"
	"aurora/services/jwt"
	"aurora/services/utils"
	"encoding/json"
	"fmt"
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

		var cacheJwt jwt.Payload
		cacheRes, err := cache.HGetAll(fmt.Sprintf("auth:%s", claims.Email))

		if err == nil && len(cacheRes) > 0 {
			_ = json.Unmarshal([]byte(cacheRes["data"]), &cacheJwt)
		}

		if cacheJwt.Email != "" {
			c.Set("user", cacheJwt)
			c.Next()
			return
		}

		// Cache miss
		var auth models.Auth
		result := db.Client.Find(&auth, "email = ?", claims.Email)

		if result.Error != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, result.Error.Error())
			c.Abort()
			return
		}

		if auth.Email == "" {
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
		serializedData, _ := json.Marshal(reqUser)

		obj := map[string]string{
			"data": string(serializedData),
		}

		_ = cache.HSet(fmt.Sprintf("auth:%s", claims.Email), obj, cache.AuthTTL)
	}
}
