package auth

import (
	"aurora/services/hash"
	"aurora/services/jwt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"unicode"
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

func customPasswordCheck(password string) bool {
	hasUpper := false
	hasLower := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}
		if unicode.IsLower(char) {
			hasLower = true
		}
		if hasLower && hasUpper {
			break
		}
	}

	return hasUpper && hasLower
}

func RegisterUser(c *gin.Context) {
	var body RegisterDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hadPassedCustomPasswordCheck := customPasswordCheck(body.Password)

	if !hadPassedCustomPasswordCheck {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation failed for password field",
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

	hashedPassword, err := hash.HashPassword(body.Password)

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

	match, err := hash.VerifyPassword(plainPassword, hashedPassword)

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

	accessToken, err := jwt.EncodeJwt(jwt.Payload{
		FullName: authResult.FullName,
		Email:    authResult.Email,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Header("x-access-token", accessToken)
	c.Header("x-refresh-token", accessToken)
	c.Status(http.StatusOK)
}
