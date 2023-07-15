package dto

type PasswordResetDto struct {
	Email       string `json:"email" binding:"required,email"`
	NewPassword string `json:"newPassword" binding:"required,min=8,max=64"`
	ResetToken  string `json:"resetToken" binding:"required"`
}
