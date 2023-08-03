package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("x-access-token")
	corsConfig.AddAllowHeaders("x-refresh-token")
	corsConfig.AllowAllOrigins = true

	return cors.New(corsConfig)
}
