package auth

type ForgotPasswordDto struct {
	Email string `json:"email" binding:"required,email"`
}
