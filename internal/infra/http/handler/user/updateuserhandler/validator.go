package updateuserhandler

import (
	"fmt"

	"github.com/cristiano-pacheco/pingo/internal/infra/validator"
)

func validateInput(in input) *validator.ValidationResult {
	v := validator.New()
	v.CheckField(validator.NotBlank(in.Name), "name", validator.NotBlankMessage)
	v.CheckField(validator.MinMaxChars(in.Name, 2, 255), "name", fmt.Sprintf(validator.MaxMaxCharsMessage, 2, 255))

	return v.Validate()
}
