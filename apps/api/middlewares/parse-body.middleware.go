package middlewares

import (
	"aurora/services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParseBody[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body T
		if err := c.ShouldBindJSON(&body); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		c.Set("body", body)

		c.Next()
	}
}
