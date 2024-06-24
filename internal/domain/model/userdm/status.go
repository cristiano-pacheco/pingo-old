package userdm

import "fmt"

type Status struct {
	value string
}

const STATUS_PENDING = "pending"
const STATUS_CONFIRMED = "confirmed"

func NewStatus(value string) (*Status, error) {
	err := validateStatus(value)
	if err != nil {
		return nil, err
	}
	return &Status{value: value}, nil
}

func (s Status) String() string {
	return s.value
}

func validateStatus(value string) error {
	if value != STATUS_PENDING && value != STATUS_CONFIRMED {
		return fmt.Errorf("the status %s value is invalid", value)
	}
	return nil
}
