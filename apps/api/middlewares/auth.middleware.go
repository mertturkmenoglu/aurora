package middlewares

import (
	authmod "aurora/routes/auth"
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

		claims, err := jwt.DecodeJwt(accessToken)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		var auth authmod.Auth
		authResult, err := auth.GetByEmail(claims.Email)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		if authResult.Email == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid credentials",
			})
			c.Abort()
			return
		}

		reqUser := jwt.Payload{
			Id:       authResult.Id,
			FullName: authResult.FullName,
			Email:    authResult.Email,
		}

		c.Set("user", reqUser)

		c.Next()
	}
}
