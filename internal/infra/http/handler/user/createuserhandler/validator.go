package createuserhandler

import (
	"fmt"

	"github.com/cristiano-pacheco/pingo/internal/infra/validator"
)

func validateInput(in input) *validator.ValidationResult {
	v := validator.New()
	v.CheckField(validator.NotBlank(in.Password), "password", validator.NotBlankMessage)
	v.CheckField(validator.NotBlank(in.Name), "name", validator.NotBlankMessage)
	v.CheckField(validator.MinMaxChars(in.Name, 2, 255), "name", fmt.Sprintf(validator.MaxMaxCharsMessage, 2, 255))
	v.CheckField(validator.MaxChars(in.Email, 255), "email", fmt.Sprintf(validator.MaxCharsMessage, 255))
	v.CheckField(validator.Matches(in.Email, validator.EmailRX), "email", validator.InvalidEmailMessage)

	return v.Validate()
}
