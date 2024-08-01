// Package emaildm contains the email domain model.
package emaildm

import (
	"net/mail"
	"strings"
)

type Email struct {
	value string
}

func New(value string) (*Email, error) {
	v, err := mail.ParseAddress(value)
	if err != nil {
		return nil, err
	}
	email := strings.ToLower(v.Address)
	return &Email{value: email}, nil
}

func (v *Email) String() string {
	return v.value
}
