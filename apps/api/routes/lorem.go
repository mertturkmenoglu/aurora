package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTestData(c *gin.Context) {
	data := []gin.H{
		{
			"name": "John Doe",
			"age":  18,
		},
		{
			"name": "Jane Doe",
			"age":  19,
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
