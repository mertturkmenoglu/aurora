package handlers

import (
	"aurora/handlers/dto"
	"aurora/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"time"
)

func Bootstrap(router *gin.Engine) {
	app := router.Group("/api/v1")

	middlewares.Limit = ratelimit.New(1000, ratelimit.Per(time.Minute))

	app.Use(middlewares.LeakBucket())

	// Auth routes
	app.POST("/auth/register", middlewares.ParseBody[dto.RegisterDto](), Register)
	app.POST("/auth/login", middlewares.ParseBody[dto.LoginDto](), Login)
	app.POST("/auth/forgot-password", middlewares.ParseBody[dto.ForgotPasswordDto](), ForgotPassword)
	app.POST("/auth/password-reset", middlewares.ParseBody[dto.PasswordResetDto](), PasswordReset)

	// User routes
	app.GET("/users/me", middlewares.IsAuth(), GetMe)

	// Product routes
	app.GET("/products/:id", GetProductById)
	app.POST("/products", middlewares.ParseBody[dto.CreateProductDto](), CreateProduct)
	app.GET("/products", GetProductByCategory)

	// Category routes
	app.POST("/categories", middlewares.ParseBody[dto.CreateCategoryDto](), CreateCategory)

	// Brand routes
	app.GET("/brands/:id", GetBrandById)
	app.POST("/brands", middlewares.ParseBody[dto.CreateBrandDto](), CreateBrand)

	// Reviews routes
	app.POST("/reviews/brands", middlewares.ParseBody[dto.CreateBrandReviewDto](), middlewares.IsAuth(), CreateBrandReview)
}
