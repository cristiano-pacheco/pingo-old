package updatecontacthandler

import (
	"fmt"

	"github.com/cristiano-pacheco/pingo/internal/domain/model/contactdm"
	"github.com/cristiano-pacheco/pingo/internal/infra/validator"
)

func validateInput(in input) *validator.ValidationResult {
	v := validator.New()
	v.CheckField(validator.NotBlank(in.Name), "name", validator.NotBlankMessage)
	v.CheckField(validator.MinMaxChars(in.Name, 2, 255), "name", fmt.Sprintf(validator.MaxMaxCharsMessage, 2, 255))

	v.CheckField(validator.NotBlank(in.ContactType), "contact_type", validator.NotBlankMessage)
	v.CheckField(validator.PermittedValue(in.ContactType, contactdm.Values()...), "contact_type", validator.PermittedValues)

	v.CheckField(validator.NotBlank(in.ContactValue), "contact_value", validator.NotBlankMessage)
	if in.ContactType == contactdm.TypeEmail {
		v.CheckField(validator.MaxChars(in.ContactValue, 255), "contact_value", fmt.Sprintf(validator.MaxCharsMessage, 255))
		v.CheckField(validator.Matches(in.ContactValue, validator.EmailRX), "contact_value", validator.InvalidEmailMessage)
	}

	if in.ContactType == contactdm.TypeSMS {
		v.CheckField(validator.Matches(in.ContactValue, validator.PhoneRX), "contact_value", validator.InvalidEmailMessage)
	}

	return v.Validate()
}
