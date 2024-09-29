package email

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/jordan-wright/email"
)

type Mailer interface {
	// SendEmail 发送邮件接口
	// subject 主题、content 内容、to 发给谁、cc 抄送给谁、bcc 密件抄送、attachFiles 附件
	SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error
}

type MailSender struct {
	smtpAuthAddress   string // 邮件服务 ip
	smtpServerAddress string // 邮件服务 ip:port
	fromName          string // 发件人
	fromEmailAddress  string // 发件人邮箱地址
	fromEmailPassword string // 发件人邮箱授权密码
}

/*
* NewMailSender 创建一个邮件发送者

* smtpAuthAddress 邮件服务ip, 如 Gmail smtpAuthAddress="smtp.gmail.com"

* smtpServerAddress 邮件服务地址，如 Gmail smtpServerAddress="smtp.gmail.com:587"

* fromName 发送邮件人名称

* fromEmailAddress 使用发邮件的邮箱地址

* fromEmailPassword 使用发邮件的邮箱密码
 */
func NewMailSender(
	smtpAuthAddress string,
	smtpServerAddress string,
	fromName string,
	fromEmailAddress string,
	fromEmailPassword string,
) Mailer {
	fmt.Printf(">>>>> smtpAuthAddress=%s, smtpServerAddress=%s\n", smtpAuthAddress, smtpServerAddress)
	return &MailSender{
		smtpAuthAddress:   smtpAuthAddress,
		smtpServerAddress: smtpServerAddress,
		fromName:          fromName,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

// SendEmail 发送邮件接口
// subject 主题、content 内容、to 发给谁、cc 抄送给谁、bcc 密件抄送、attachFiles 附件
func (sender *MailSender) SendEmail(subject string, content string, to []string, cc []string, bcc []string, attachFiles []string) error {
	fmt.Printf(">>>>> 1111 smtpAuthAddress=%s, smtpServerAddress=%s\n", sender.smtpAuthAddress, sender.smtpServerAddress)

	mail := email.NewEmail()
	mail.From = sender.fromName + "<" + sender.fromEmailAddress + ">"
	mail.Subject = subject
	mail.HTML = []byte(content)
	mail.To = to
	mail.Cc = cc
	mail.Bcc = bcc

	for _, f := range attachFiles {
		_, err := mail.AttachFile(f)
		if err != nil {
			return fmt.Errorf("加载附近出错,请确保路径正确 %s: %w", f, err)
		}
	}

	var auth smtp.Auth
	if os.Getenv("ENV") == "prod" {
		// NOTE: 会走TLS加密
		// identity 通常是空字符串
		auth = smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, sender.smtpAuthAddress)
	} else {
		// 非加密
		auth = smtp.CRAMMD5Auth(sender.fromName, sender.fromEmailPassword)
	}

	return mail.Send(sender.smtpServerAddress, auth)
}
