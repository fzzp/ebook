package email

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestSendEmail(t *testing.T) {
	godotenv.Load("../../.env")
	sender := NewMailSender(
		os.Getenv("SMTP_AUTH_ADDRESS"),
		os.Getenv("SMTP_SERVER_ADDRESS"),
		os.Getenv("MAIL_SENDER_NAME"),
		os.Getenv("MAIL_SENDER_ADDRESS"),
		os.Getenv("MAIL_SENDER_PASSWORD"),
	)

	subject := "Test Send Email"
	content := `
		<h1>您好，欢迎注册 Ebook</h1>
	`
	to := []string{os.Getenv("TEST_EMAIL_TO")}
	attchFiles := []string{"../../README.md"}

	err := sender.SendEmail(subject, content, to, nil, nil, attchFiles)
	if err != nil {
		t.Error(err)
	}
}
