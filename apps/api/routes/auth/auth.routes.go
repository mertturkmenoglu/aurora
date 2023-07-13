package auth

import (
	"aurora/services/aws/models"
	"aurora/services/cache"
	"aurora/services/hash"
	"aurora/services/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
	"unicode"
)

func doesUserExist(email string) (bool, error) {
	var auth models.Auth
	authResult, err := auth.GetByEmail(email)

	if err != nil {
		return false, err
	}

	if (models.Auth{}) == *authResult {
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

func generateRandomShortId() string {
	sid := uuid.NewString()
	return strings.Split(sid, "-")[0]
}

func RegisterUser(c *gin.Context) {
	body := c.MustGet("body").(RegisterDto)

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

	hashedPassword, err := hash.StringHash(body.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := uuid.NewString()
	auth := models.Auth{
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
	body := c.MustGet("body").(LoginDto)

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

	var auth models.Auth
	authResult, err := auth.GetByEmail(body.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	plainPassword := body.Password
	hashedPassword := authResult.Password

	match, err := hash.VerifyHash(plainPassword, hashedPassword)

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
		Id:       authResult.Id,
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

func ForgotPassword(c *gin.Context) {
	body := c.MustGet("body").(ForgotPasswordDto)

	userExistsResult, err := doesUserExist(body.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !userExistsResult {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User does not exist",
		})
		return
	}

	// Generate random short id
	// Then create a formatted string with email and sid
	// Then hash the formatted string
	// Then save the hashed string to Redis
	// Then send the sid to the user's email
	sid := generateRandomShortId()
	formattedEmailSid := fmt.Sprintf("%s:%s", body.Email, sid)

	hashed, err := hash.StringHash(formattedEmailSid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	key := fmt.Sprintf("forgot-password:%s", body.Email)
	err = cache.Set(key, hashed, time.Minute*15)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO: Send email to user's email
	fmt.Println("sid", sid)

	c.Status(http.StatusOK)
}

func PasswordReset(c *gin.Context) {
	body := c.MustGet("body").(PasswordResetDto)

	// Check if new password passes custom validation
	// If it doesn't, then return an error
	hadPassedCustomPasswordCheck := customPasswordCheck(body.NewPassword)

	if !hadPassedCustomPasswordCheck {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation failed for password field",
		})
		return
	}

	// Check if user exists
	userExistsResult, err := doesUserExist(body.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !userExistsResult {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User does not exist",
		})
		return
	}

	// Get the hashed string from Redis
	// Then compare the hashed string with the hashed string from the request body
	// If they match, then update the user's password
	// If they don't match, then return an error
	key := fmt.Sprintf("forgot-password:%s", body.Email)
	hashed, err := cache.Get(key)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	formattedEmailSid := fmt.Sprintf("%s:%s", body.Email, body.ResetToken)
	match, err := hash.VerifyHash(formattedEmailSid, hashed)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid reset token",
		})
		return
	}

	// Update the user's password
	var auth models.Auth
	authResult, err := auth.GetByEmail(body.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashedPassword, err := hash.StringHash(body.NewPassword)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	authResult.Password = hashedPassword
	_, err = authResult.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Delete the hashed string from Redis
	_ = cache.Del(key)

	c.Status(http.StatusOK)
}
