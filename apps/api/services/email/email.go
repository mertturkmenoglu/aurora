package email

import (
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

type Config struct {
	FromName     string
	FromEmail    string
	SmtpEmail    string
	SmtpPassword string
}

type WithTemplateConfig[T Payload] struct {
	TemplatePath string
	Data         T
	To           string
	Subject      string
}

func getEmailConfig() Config {
	fromName, fromNameOk := os.LookupEnv("EMAIL_FROM_NAME")
	fromEmail, fromEmailOk := os.LookupEnv("EMAIL_FROM_EMAIL")
	smtpEmail, smtpEmailOk := os.LookupEnv("SMTP_EMAIL")
	smtpPassword, smtpPasswordOk := os.LookupEnv("SMTP_PASSWORD")

	if !fromNameOk || !fromEmailOk || !smtpEmailOk || !smtpPasswordOk {
		panic("Email environment variables are not set")
	}

	return Config{
		FromName:     fromName,
		FromEmail:    fromEmail,
		SmtpEmail:    smtpEmail,
		SmtpPassword: smtpPassword,
	}
}

func SendEmail(to string, subject string, text string) error {
	cfg := getEmailConfig()
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", cfg.FromName, cfg.FromEmail)
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)

	auth := smtp.PlainAuth("", cfg.SmtpEmail, cfg.SmtpPassword, "smtp.gmail.com")

	err := e.Send("smtp.gmail.com:587", auth)

	if err != nil {
		log.Println("Error sending email: ", err)
		return err
	}

	return nil
}

func SendEmailWithTemplate[T Payload](templateConfig WithTemplateConfig[T]) error {
	cfg := getEmailConfig()

	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", cfg.FromName, cfg.FromEmail)
	e.To = []string{templateConfig.To}

	t, err := template.ParseFiles(templateConfig.TemplatePath)

	if err != nil {
		return err
	}

	var body bytes.Buffer

	err = t.Execute(&body, templateConfig.Data)

	if err != nil {
		return err
	}

	e.Subject = templateConfig.Subject
	e.HTML = body.Bytes()

	auth := smtp.PlainAuth("", cfg.SmtpEmail, cfg.SmtpPassword, "smtp.gmail.com")

	err = e.Send("smtp.gmail.com:587", auth)

	return err
}
