// Package errordm contains the Error domain model.
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

func (v *Error) Add(key string, err error) {
	if err == nil {
		return
	}
	if _, exists := v.errors[key]; !exists {
		v.errors[key] = err.Error()
	}
}

func (v *Error) String() string {
	return v.toJSON()
}

func (v *Error) Error() error {
	return errors.New(string(v.toJSON()))
}

func (v *Error) toJSON() string {
	if len(v.errors) == 0 {
		return ""
	}

	jsonData, err := json.Marshal(v.errors)
	if err != nil {
		return `{"error": "failed to marshal errors"}`
	}

	return string(jsonData)
}
