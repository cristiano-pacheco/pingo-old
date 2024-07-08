// Package mailertemplategw contains the MailerTemplateGateway interface.
package mailertemplategw

type MailerTemplateGateway interface {
	CompileTemplate(templateName string, templateData any) (string, error)
	CompileBlankTemplate(templateName string, data any) (string, error)
}
