package phonedm

import (
	"fmt"
	"regexp"
)

type Phone struct {
	value string
}

func New(value string) (*Phone, error) {
	err := validate(value)
	if err != nil {
		return nil, err
	}
	return &Phone{value}, nil
}

func (p *Phone) String() string {
	return p.value
}

func validate(value string) error {
	e164Regex := regexp.MustCompile(`^\+\d{1,3}\d{7,14}$`)
	result := e164Regex.MatchString(value)

	if !result {
		return fmt.Errorf("the phone number %s is invalid", value)
	}

	return nil
}
