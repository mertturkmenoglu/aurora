package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/handlers/dto"
	"aurora/services/jwt"
	"aurora/services/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

func GetMyCart(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)

	var user models.User

	res := db.Client.
		Where("email = ?", reqUser.Email).
		First(&user)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	var cart models.Cart

	res = db.Client.
		Preload(clause.Associations).
		Where("user_id = ?", user.Id).
		First(&cart)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.JSON(200, gin.H{
		"data": cart,
	})
}

func AddToCart(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)
	body := c.MustGet("body").(dto.AddToCartDto)

	// Get user by request user email
	var user models.User

	res := db.Client.
		Where("email = ?", reqUser.Email).
		First(&user)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	// Get cart by user id
	var cart models.Cart

	res = db.Client.
		Preload(clause.Associations).
		Where("user_id = ?", user.Id).
		First(&cart)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	// Get product by product id
	var product models.Product

	res = db.Client.
		Preload(clause.Associations).
		Where("id = ?", body.ProductId).
		First(&product)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	// Check if product is already in cart
	// By default, assume it's already in cart to prevent double adding
	isProductAlreadyInCart := true
	var cartItem models.CartItem

	res = db.Client.
		Where("cart_id = ? AND product_id = ?", cart.Id, product.Id).
		First(&cartItem)

	if res.Error != nil {
		// If cart item is not found, then it's not in cart
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			isProductAlreadyInCart = false
		} else {
			// If there's an error other than record not found, then return error
			utils.HandleDatabaseError(c, res.Error)
			return
		}
	}

	// If product is already in cart, then return error
	if isProductAlreadyInCart {
		utils.ErrorResponse(c, 400, "Product is already in cart")
		return
	}

	newCartItem := models.CartItem{
		CartId:    cart.Id,
		ProductId: product.Id,
		Quantity:  body.Quantity,
	}

	res = db.Client.
		Preload(clause.Associations).
		Create(&newCartItem)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	newCartItem.Product = product

	c.JSON(http.StatusCreated, gin.H{
		"data": newCartItem,
	})
}

func RemoveFromCart(c *gin.Context) {

}

func RemoveAllFromCart(c *gin.Context) {

}

func UpdateCartItem(c *gin.Context) {

}
