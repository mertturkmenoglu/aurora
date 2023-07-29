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
	app.POST("/auth/password/forgot", middlewares.ParseBody[dto.ForgotPasswordDto](), ForgotPassword)
	app.POST("/auth/password/reset", middlewares.ParseBody[dto.ResetPasswordDto](), ResetPassword)
	app.PUT("/auth/password/change", middlewares.IsAuth(), ChangePassword)

	// User routes
	app.GET("/users/me", middlewares.IsAuth(), GetMe)
	app.PUT("/users/me", middlewares.ParseBody[dto.UpdateUserDto](), middlewares.IsAuth(), UpdateMe)
	app.GET("/users/me/addresses", middlewares.IsAuth(), GetMyAddresses)
	app.POST("/users/me/addresses", middlewares.IsAuth(), AddAddress)
	app.PUT("/users/me/addresses/:id", middlewares.IsAuth(), UpdateAddress)
	app.DELETE("/users/me/addresses/:id", middlewares.IsAuth(), DeleteAddress)
	app.PUT("/users/me/ad-preferences", middlewares.IsAuth(), UpdateAdPreferences)

	// Product routes
	app.GET("/products/all", GetAllProducts)
	app.GET("/products/featured", GetFeaturedProducts)
	app.GET("/products/new", GetNewProducts)
	app.GET("/products/sale", GetSaleProducts)
	app.GET("/products/popular", GetPopularProducts)
	app.GET("/products/free-shipping", GetFreeShippingProducts)
	app.GET("/products/:id", GetProductById)
	app.POST("/products/:id/styles", middlewares.ParseBody[dto.AddProductStylesDto](), AddProductStyles)
	app.POST("/products/:id/sizes", middlewares.ParseBody[dto.AddProductSizesDto](), AddProductSizes)
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

	// Favorites routes
	app.GET("/favorites", middlewares.IsAuth(), GetMyFavorites)
	app.POST("/favorites", middlewares.ParseBody[dto.AddFavoriteDto](), middlewares.IsAuth(), AddFavorite)
	app.DELETE("/favorites/:id", middlewares.IsAuth(), DeleteFavorite)
	app.DELETE("/favorites", middlewares.IsAuth(), DeleteAllFavorites)

	// Search routes
	app.GET("/search", SearchProducts)

	// Cart routes
	app.GET("/cart", middlewares.IsAuth(), GetMyCart)
	app.POST("/cart", middlewares.ParseBody[dto.AddToCartDto](), middlewares.IsAuth(), AddToCart)
	app.DELETE("/cart/:id", middlewares.IsAuth(), RemoveFromCart)
	app.DELETE("/cart", middlewares.IsAuth(), RemoveAllFromCart)
	app.PUT("/cart/:id", middlewares.ParseBody[dto.UpdateCartItemDto](), middlewares.IsAuth(), UpdateCartItem)
}
