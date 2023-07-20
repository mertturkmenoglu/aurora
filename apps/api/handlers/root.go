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
	app.PUT("/auth/password-change", func(c *gin.Context) {}) // TODO: Change password

	// User routes
	app.GET("/users/me", middlewares.IsAuth(), GetMe)
	app.PUT("/users/me", middlewares.ParseBody[dto.UpdateUserDto](), middlewares.IsAuth(), UpdateMe)
	app.GET("/users/me/addresses", middlewares.IsAuth(), GetMyAddresses)
	app.POST("/users/me/addresses", func(c *gin.Context) {})       // TODO: Create an address
	app.PUT("/users/me/addresses/:id", func(c *gin.Context) {})    // TODO: Update an address
	app.DELETE("/users/me/addresses/:id", func(c *gin.Context) {}) // TODO: Delete an address
	app.PUT("/users/me/ad-preferences", func(c *gin.Context) {})   // TODO: Update my ad preferences

	// Product routes
	app.GET("/products/:id", GetProductById)
	app.POST("/products", middlewares.ParseBody[dto.CreateProductDto](), CreateProduct)
	app.GET("/products", GetProductByCategory)

	// Category routes
	app.GET("/categories", GetCategories)
	app.POST("/categories", middlewares.ParseBody[dto.CreateCategoryDto](), CreateCategory)

	// Brand routes
	app.GET("/brands/:id", GetBrandById)
	app.POST("/brands", middlewares.ParseBody[dto.CreateBrandDto](), CreateBrand)

	// Reviews routes
	app.POST("/reviews/brands", middlewares.ParseBody[dto.CreateBrandReviewDto](), middlewares.IsAuth(), CreateBrandReview)
	app.POST("/reviews/products", middlewares.ParseBody[dto.CreateProductReviewDto](), middlewares.IsAuth(), CreateProductReview)
	app.GET("/reviews/my-reviews/brands", middlewares.IsAuth(), GetMyBrandReviews)
	app.GET("/reviews/my-reviews/products", middlewares.IsAuth(), GetMyProductReviews)
	app.GET("/reviews/brands/:id", GetBrandReview)
	app.GET("/reviews/products/:id", GetProductReview)
	app.GET("/reviews/brands", GetBrandReviews)
	app.GET("/reviews/products", GetProductReviews)
	app.DELETE("/reviews/brands/:id", middlewares.IsAuth(), DeleteBrandReview)
	app.DELETE("/reviews/products/:id", middlewares.IsAuth(), DeleteProductReview)
}
