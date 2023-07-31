package dto

type ChangePasswordDto struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=8,max=64"`
}
