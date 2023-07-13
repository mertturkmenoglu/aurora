package routes

import (
	"aurora/middlewares"
	"aurora/routes/auth"
	"aurora/routes/products"
	"aurora/routes/users"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"time"
)

func Bootstrap(router *gin.Engine) {
	app := router.Group("/v1")

	middlewares.Limit = ratelimit.New(1000, ratelimit.Per(time.Minute))

	app.Use(middlewares.LeakBucket())

	// Auth routes
	app.POST("/auth/register", middlewares.ParseBody[auth.RegisterDto](), auth.RegisterUser)
	app.POST("/auth/login", middlewares.ParseBody[auth.LoginDto](), auth.LoginUser)
	app.POST("/auth/forgot-password", middlewares.ParseBody[auth.ForgotPasswordDto](), auth.ForgotPassword)
	app.POST("/auth/password-reset", middlewares.ParseBody[auth.PasswordResetDto](), auth.PasswordReset)

	// User routes
	app.GET("/users/:email", middlewares.IsAuth(), users.GetUserById)

	// Product routes
	app.GET("/products/:id", products.GetProductById)
}
