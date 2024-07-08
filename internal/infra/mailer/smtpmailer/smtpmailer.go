// Package smtpmailer contains the implementation of mailergateway.
package smtpmailer

import (
	"github.com/cristiano-pacheco/pingo/internal/application/gateway/mailergw"
	"github.com/go-mail/mail/v2"
)

type SMTPMailer struct {
	dialer *mail.Dialer
	sender string
}

func New(dialer *mail.Dialer, sender string) *SMTPMailer {
	return &SMTPMailer{dialer: dialer, sender: sender}
}

func (m *SMTPMailer) Send(md *mailergw.MailData) error {
	msg := mail.NewMessage()
	msg.SetHeader("To", md.ToEmail)
	msg.SetHeader("From", m.sender)
	msg.SetHeader("Subject", md.Subject)
	msg.SetBody("text/html", md.Content)

	err := m.dialer.DialAndSend(msg)
	if err != nil {
		return err
	}

	return nil
}
