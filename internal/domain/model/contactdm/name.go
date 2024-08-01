package contactdm

import (
	"fmt"
	"unicode/utf8"
)

type Name struct {
	value string
}

func NewName(value string) (*Name, error) {
	err := validateName(value)
	if err != nil {
		return nil, err
	}
	return &Name{value: value}, nil
}

func (v Name) String() string {
	return v.value
}

func validateName(value string) error {
	total := utf8.RuneCountInString(value)
	if total < 3 || total > 255 {
		return fmt.Errorf("the name must contain between 2 and 255 characters")
	}
	return nil
}
