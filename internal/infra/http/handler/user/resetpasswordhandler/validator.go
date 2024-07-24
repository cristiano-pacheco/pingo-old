package resetpasswordhandler

import (
	"github.com/cristiano-pacheco/pingo/internal/infra/validator"
)

func validateInput(in input) *validator.ValidationResult {
	v := validator.New()
	v.CheckField(validator.NotBlank(in.ID), "id", validator.NotBlankMessage)
	v.CheckField(validator.NotBlank(in.Token), "token", validator.NotBlankMessage)
	v.CheckField(validator.NotBlank(in.Password), "password", validator.NotBlankMessage)

	return v.Validate()
}
