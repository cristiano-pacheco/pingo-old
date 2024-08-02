// Package contactrepo contains the contact repository.
package contactrepo

import (
	"github.com/cristiano-pacheco/pingo/internal/domain/model/contactdm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
)

type ContactRepository interface {
	Create(contact contactdm.Contact) error
	Update(contact contactdm.Contact) error
	Delete(contact contactdm.Contact) error
	FindByIDAndUserID(id, userID identitydm.ID) (*contactdm.Contact, error)
	FindListByUserID(userID identitydm.ID) ([]*contactdm.Contact, error)
}
