package userdm

import "fmt"

type Status struct {
	value string
}

const StatusPending = "pending"
const StatusActivated = "activated"
const StatusDisabled = "disabled"

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
	if value != StatusPending && value != StatusActivated && value != StatusDisabled {
		return fmt.Errorf("the status %s value is invalid", value)
	}
	return nil
}
