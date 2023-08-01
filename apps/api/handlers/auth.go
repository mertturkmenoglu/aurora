package handlers

import (
	"aurora/db"
	"aurora/db/models"
	"aurora/handlers/dto"
	"aurora/services/cache"
	"aurora/services/email"
	"aurora/services/hash"
	"aurora/services/jwt"
	"aurora/services/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	body := c.MustGet("body").(dto.RegisterDto)

	hadPassedCustomPasswordCheck := utils.CustomPasswordCheck(body.Password)

	if !hadPassedCustomPasswordCheck {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed for the password field")
		return
	}

	userExists, err := utils.DoesUserExist(body.Email)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if userExists {
		utils.ErrorResponse(c, http.StatusBadRequest, "User already exists")
		return
	}

	hashed, err := hash.StringHash(body.Password)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	auth := models.Auth{
		FullName: body.FullName,
		Email:    body.Email,
		Password: hashed,
	}

	res := db.Client.Create(&auth)

	if res.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, res.Error.Error())
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
		Addresses: []models.Address{},
		Phone:     "",
	}

	res = db.Client.Create(&user)

	if res.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, res.Error.Error())
		return
	}

	res = db.Client.Create(&models.Cart{
		UserId: user.Id,
	})

	if res.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, res.Error.Error())
		return
	}

	c.Status(http.StatusCreated)
}

func Login(c *gin.Context) {
	body := c.MustGet("body").(dto.LoginDto)

	userExists, err := utils.DoesUserExist(body.Email)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !userExists {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid credentials")
		return
	}

	var auth models.Auth

	res := db.Client.First(&auth, "email = ?", body.Email)

	if res.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, res.Error.Error())
		return
	}

	plain := body.Password
	hashed := auth.Password

	match, err := hash.VerifyHash(plain, hashed)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !match {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid credentials")
		return
	}

	var user models.User

	res = db.Client.First(&user, "email = ?", body.Email)

	if res.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, res.Error.Error())
		return
	}

	token, err := jwt.EncodeJwt(jwt.Payload{
		UserId:   user.Id.String(),
		Email:    user.Email,
		FullName: user.FullName,
	})

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"accessToken":  token,
			"refreshToken": token,
		},
	})
}

func ForgotPassword(c *gin.Context) {
	body := c.MustGet("body").(dto.ForgotPasswordDto)

	userExists, err := utils.DoesUserExist(body.Email)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !userExists {
		utils.ErrorResponse(c, http.StatusBadRequest, "User does not exist")
		return
	}

	sid := utils.GenerateRandomShortId()
	formatted := fmt.Sprintf("%s:%s", body.Email, sid)
	hashed, err := hash.StringHash(formatted)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	key := cache.GetFormattedKey(cache.ForgotPasswordKeyFormat, body.Email)
	err = cache.Set(key, hashed, cache.ForgotPasswordTTL)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: Delegate this to a worker
	err = email.SendEmailWithTemplate(email.WithTemplateConfig[email.ForgotPasswordPayload]{
		To:           body.Email,
		TemplatePath: "templates/forgot-password.html",
		Subject:      "Reset your password",
		Data: email.ForgotPasswordPayload{
			Code: sid,
		},
	})

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func ResetPassword(c *gin.Context) {
	body := c.MustGet("body").(dto.ResetPasswordDto)

	hadPassedCustomPasswordCheck := utils.CustomPasswordCheck(body.NewPassword)

	if !hadPassedCustomPasswordCheck {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed for the password field")
		return
	}

	userExists, err := utils.DoesUserExist(body.Email)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !userExists {
		utils.ErrorResponse(c, http.StatusBadRequest, "User does not exist")
		return
	}

	key := cache.GetFormattedKey(cache.ForgotPasswordKeyFormat, body.Email)
	hashed, err := cache.Get(key)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	match, err := hash.VerifyHash(body.ResetToken, hashed)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !match {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid reset token")
		return
	}

	var auth models.Auth
	res := db.Client.First(&auth, "email = ?", body.Email)

	if res.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, res.Error.Error())
		return
	}

	hashedPassword, err := hash.StringHash(body.NewPassword)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	auth.Password = hashedPassword
	res = db.Client.Save(&auth)

	if res.Error != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, res.Error.Error())
		return
	}

	_ = cache.Del(key)

	c.Status(http.StatusOK)
}

func ChangePassword(c *gin.Context) {
	reqUser := c.MustGet("user").(jwt.Payload)
	body := c.MustGet("body").(dto.ChangePasswordDto)

	var auth models.Auth

	res := db.Client.First(&auth, "email = ?", reqUser.Email)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	match, err := hash.VerifyHash(body.OldPassword, auth.Password)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !match {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid credentials")
		return
	}

	hadPassedCustomPasswordCheck := utils.CustomPasswordCheck(body.NewPassword)

	if !hadPassedCustomPasswordCheck {
		utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed for the password field")
		return
	}

	hashedPassword, err := hash.StringHash(body.NewPassword)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	auth.Password = hashedPassword

	res = db.Client.Save(&auth)

	if res.Error != nil {
		utils.HandleDatabaseError(c, res.Error)
		return
	}

	c.Status(http.StatusOK)
}
