package errordm

import (
	"encoding/json"
	"errors"
)

type Error struct {
	errors map[string]string
}

func New() *Error {
	return &Error{
		errors: make(map[string]string),
	}
}

func (v *Error) Add(key, value string) {
	if _, exists := v.errors[key]; !exists {
		v.errors[key] = value
	}
}

func (v *Error) String() string {
	return v.toJson()
}

func (v *Error) Error() error {
	return errors.New(string(v.toJson()))
}

func (v *Error) toJson() string {
	if len(v.errors) == 0 {
		return ""
	}

	jsonData, err := json.Marshal(v.errors)
	if err != nil {
		return `{"error": "failed to marshal errors"}`
	}

	return string(jsonData)
}
