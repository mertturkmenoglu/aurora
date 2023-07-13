package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParseBody[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body T
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("body", body)

		c.Next()
	}
}
