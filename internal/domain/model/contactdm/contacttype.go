package contactdm

import "fmt"

type ContactType struct {
	value string
}

const TypeEmail = "email"
const TypeSMS = "sms"
const TypeSlack = "slack"

func NewContactType(value string) (*ContactType, error) {
	err := validateContactType(value)
	if err != nil {
		return nil, err
	}
	return &ContactType{value: value}, nil
}

func (s ContactType) String() string {
	return s.value
}

func validateContactType(value string) error {
	if value != TypeEmail && value != TypeSMS && value != TypeSlack {
		return fmt.Errorf("the contact type %s value is invalid", value)
	}
	return nil
}
