package middlewares

import (
	"aurora/services/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("x-access-token")

		if accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Missing access token",
			})
			c.Abort()
			return
		}

		_, err := jwt.DecodeJwt(accessToken)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		// TODO: Verify user

		c.Next()
	}
}
