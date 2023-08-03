package tasks

type EmailForgotPasswordPayload struct {
	Email string
	Code  string
}
