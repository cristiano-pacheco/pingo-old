package contactdm

import (
	"fmt"

	"github.com/cristiano-pacheco/pingo/internal/domain/model/emaildm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/phonedm"
)

type ContactData struct {
	contactType ContactType
	value       string
}

func NewContactData(contactType, value string) (*ContactData, error) {
	ctype, err := NewContactType(contactType)
	if err != nil {
		return nil, err
	}

	if ctype.String() == TypeEmail {
		emailVo, err := emaildm.New(value)
		if err != nil {
			return nil, err
		}
		return &ContactData{*ctype, emailVo.String()}, nil
	}

	if ctype.String() == TypeSMS {
		phoneVo, err := phonedm.New(value)
		if err != nil {
			return nil, err
		}
		return &ContactData{*ctype, phoneVo.String()}, nil
	}

	return nil, fmt.Errorf("the contact data type: %s value: %s is invalid", ctype.String(), value)
}

func (cd *ContactData) ContactType() string {
	return cd.contactType.String()
}

func (cd *ContactData) ContactValue() string {
	return cd.value
}
