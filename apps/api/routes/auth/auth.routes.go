package auth

import (
	"aurora/services/cache"
	"aurora/services/db"
	"aurora/services/db/models"
	"aurora/services/hash"
	"aurora/services/jwt"
	"aurora/services/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterUser(c *gin.Context) {
	body := c.MustGet("body").(RegisterDto)

	// Password validation
	// Password must have at least one uppercase and one lowercase letter
	// Other password validations are handled with the binding tag
	hadPassedCustomPasswordCheck := customPasswordCheck(body.Password)

	if !hadPassedCustomPasswordCheck {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed for password field")
		return
	}

	userExistsResult, err := doesUserExist(body.Email)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userExistsResult {
		utils.ErrorResponse(c, http.StatusBadRequest, "User already exists")
		return
	}

	hashedPassword, err := hash.StringHash(body.Password)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	auth := models.Auth{
		FullName: body.FullName,
		Email:    body.Email,
		Password: hashedPassword,
	}

	result := db.Client.Create(&auth)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	user := models.User{
		FullName: body.FullName,
		Email:    body.Email,
		AdPreference: models.AdPreference{
			Email: body.HasAcceptedEmailCampaign,
			Sms:   false,
			Phone: false,
		},
		Addresses: make([]models.Address, 0),
		Phone:     "",
	}

	result = db.Client.Create(&user)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func LoginUser(c *gin.Context) {
	body := c.MustGet("body").(LoginDto)

	userExistsResult, err := doesUserExist(body.Email)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !userExistsResult {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid credentials")
		return
	}

	var auth models.Auth

	result := db.Client.First(&auth, "email = ?", body.Email)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	plainPassword := body.Password
	hashedPassword := auth.Password

	match, err := hash.VerifyHash(plainPassword, hashedPassword)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !match {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid credentials")
		return
	}

	accessToken, err := jwt.EncodeJwt(jwt.Payload{
		Id:       auth.ID.String(),
		FullName: auth.FullName,
		Email:    auth.Email,
	})

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
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
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !userExistsResult {
		utils.ErrorResponse(c, http.StatusBadRequest, "User does not exist")
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
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	key := fmt.Sprintf("forgot-password:%s", body.Email)
	err = cache.Set(key, hashed, time.Minute*15)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
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
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed for password field")
		return
	}

	// Check if user exists
	userExistsResult, err := doesUserExist(body.Email)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !userExistsResult {
		utils.ErrorResponse(c, http.StatusBadRequest, "User does not exist")
		return
	}

	// Get the hashed string from Redis
	// Then compare the hashed string with the hashed string from the request body
	// If they match, then update the user's password
	// If they don't match, then return an error
	key := fmt.Sprintf("forgot-password:%s", body.Email)
	hashed, err := cache.Get(key)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	formattedEmailSid := fmt.Sprintf("%s:%s", body.Email, body.ResetToken)
	match, err := hash.VerifyHash(formattedEmailSid, hashed)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !match {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid reset token")
		return
	}

	// Update the user's password
	var auth models.Auth
	result := db.Client.First(&auth, "email = ?", body.Email)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	hashedPassword, err := hash.StringHash(body.NewPassword)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	auth.Password = hashedPassword

	result = db.Client.Save(&auth)

	if result.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	// Delete the hashed string from Redis
	_ = cache.Del(key)

	c.Status(http.StatusOK)
}
