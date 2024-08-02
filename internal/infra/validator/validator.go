// Package validator contains the validator a validation rules
package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	PhoneRX = regexp.MustCompile(`^\+\d{1,3}\d{7,14}$`)
)

var (
	NotBlankMessage     = "must be provided"
	MinCharsMessage     = "must be at least %d characters long"
	MaxCharsMessage     = "must not be more than %d characters long"
	MaxMaxCharsMessage  = "must be between %d and %d characters"
	InvalidEmailMessage = "must be a valid email address"
	PermittedValues     = "must be a permitted value"
)

// Validator type which contains a map of validation errors for our form fields.
type Validator struct {
	errors []*error
}

func New() *Validator {
	errors := make([]*error, 0, 1)
	return &Validator{errors: errors}
}

type error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationResult struct {
	Errors  []*error `json:"errors"`
	IsValid bool     `json:"isValid"`
}

func (v *Validator) Validate() *ValidationResult {
	return &ValidationResult{
		Errors:  v.errors,
		IsValid: len(v.errors) == 0,
	}
}

// CheckField adds an error message to the Errors map only if a validation check is not 'ok'.
func (v *Validator) CheckField(ok bool, field, message string) {
	if !ok {
		v.addFieldError(field, message)
	}
}

// addFieldError adds an error message to the Errors map (so long as no
// entry already exists for the given field).
func (v *Validator) addFieldError(field, message string) {
	v.errors = append(v.errors, &error{Field: field, Message: message})
}

// NotBlank returns true if a value is not an empty string.
func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars returns true if a value contains no more than n characters.
func MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// MinChars returns true if a value contains less than n characters.
func MinChars(value string, n int) bool {
	return utf8.RuneCountInString(value) >= n
}

// MinMaxChars returns true if a value contains less than min or more than max characters.
func MinMaxChars(value string, min, max int) bool {
	totalRunes := utf8.RuneCountInString(value)
	return !(totalRunes < min || totalRunes > max)
}

// GreaterThanZero returns true if a value is greater than zero
func GreaterThanZero(value int) bool {
	return value > 0
}

// PermittedInt returns true if a value is in a list of permitted integers.
func PermittedInt(value int, permittedValues ...int) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func PermittedValue[T comparable](value T, permittedValues ...T) bool {
	for i := range permittedValues {
		if value == permittedValues[i] {
			return true
		}
	}
	return false
}
