package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func doesUserExist(email string) (bool, error) {
	var auth Auth
	authResult, err := auth.GetByEmail(email)

	if err != nil {
		return false, err
	}

	if (Auth{}) == *authResult {
		return false, nil
	}

	return true, nil
}

func RegisterUser(c *gin.Context) {
	var body RegisterDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userExistsResult, err := doesUserExist(body.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if userExistsResult {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
		return
	}

	hashedPassword, err := hashPassword(body.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := uuid.NewString()
	auth := Auth{
		Id:       id,
		FullName: body.FullName,
		Email:    body.Email,
		Password: hashedPassword,
	}

	result, err := auth.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
	})
}

func LoginUser(c *gin.Context) {
	var body LoginDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userExistsResult, err := doesUserExist(body.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !userExistsResult {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	var auth Auth
	authResult, err := auth.GetByEmail(body.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	plainPassword := body.Password
	hashedPassword := authResult.Password

	match, err := verifyPassword(plainPassword, hashedPassword)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": match,
	})
}
