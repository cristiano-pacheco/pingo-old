// Package mailergw contains the MailerGateway interface.
package mailergw

type MailerGateway interface {
	Send(md *MailData) error
}
