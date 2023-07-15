package dto

type RegisterDto struct {
	FullName                 string `json:"fullName" binding:"required"`
	Email                    string `json:"email" binding:"required,email"`
	Password                 string `json:"password" binding:"required,min=8,max=64"`
	VerifyPassword           string `json:"verifyPassword" binding:"required,eqfield=Password"`
	HasAcceptedEmailCampaign bool   `json:"hasAcceptedEmailCampaign" validate:"exists"`
}
