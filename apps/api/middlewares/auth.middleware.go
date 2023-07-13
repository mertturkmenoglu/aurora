package middlewares

import (
	"aurora/services/aws/models"
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

		var auth models.Auth
		authResult, err := auth.GetByEmail(claims.Email)

		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		if authResult.Email == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
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
