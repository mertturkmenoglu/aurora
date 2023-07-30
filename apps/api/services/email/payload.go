package email

type Payload interface {
	WelcomePayload | ForgotPasswordPayload
}
