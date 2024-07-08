// Package mailertemplate contains the implementation and the email templates.
package mailertemplate

import (
	"bytes"
	"embed"
	"html/template"
)

//go:embed "templates"
var templateFS embed.FS

type MailerTemplate struct {
}

func (mt MailerTemplate) CompileTemplate(templateName string, data any) (string, error) {
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/layout/default.gohtml", "templates/"+templateName)
	if err != nil {
		return "", err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return "", err
	}

	return htmlBody.String(), nil
}

func (mt MailerTemplate) CompileBlankTemplate(templateName string, data any) (string, error) {
	tmpl, err := template.New("email").ParseFS(templateFS, "templates/"+templateName)
	if err != nil {
		return "", err
	}

	htmlBody := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(htmlBody, "htmlBody", data)
	if err != nil {
		return "", err
	}

	return htmlBody.String(), nil
}
