package tasks

import (
	"aurora/services/email"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
)

func NewEmailForgotPasswordTask(email string, code string) (*asynq.Task, error) {
	payload, err := json.Marshal(EmailForgotPasswordPayload{
		Email: email,
		Code:  code,
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeEmailForgotPassword, payload), nil
}

func HandleEmailForgotPasswordTask(_ context.Context, t *asynq.Task) error {
	var p EmailForgotPasswordPayload

	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	err := email.SendEmailWithTemplate(email.WithTemplateConfig[email.ForgotPasswordPayload]{
		To:           p.Email,
		TemplatePath: "templates/forgot-password.html",
		Subject:      "Reset your password",
		Data: email.ForgotPasswordPayload{
			Code: p.Code,
		},
	})

	return err
}
