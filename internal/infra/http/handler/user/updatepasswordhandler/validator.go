package updatepasswordhandler

import (
	"github.com/cristiano-pacheco/pingo/internal/infra/validator"
)

func validateInput(in input) *validator.ValidationResult {
	v := validator.New()
	v.CheckField(validator.NotBlank(in.CurrentPassword), "current_password", validator.NotBlankMessage)
	v.CheckField(validator.NotBlank(in.NewPassword), "new_password", validator.NotBlankMessage)

	return v.Validate()
}
