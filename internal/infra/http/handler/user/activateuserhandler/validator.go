package activateuserhandler

import (
	"github.com/cristiano-pacheco/pingo/internal/infra/validator"
)

func validateInput(in input) *validator.ValidationResult {
	v := validator.New()
	v.CheckField(validator.NotBlank(in.ID), "id", validator.NotBlankMessage)
	v.CheckField(validator.NotBlank(in.Token), "token", validator.NotBlankMessage)

	return v.Validate()
}
