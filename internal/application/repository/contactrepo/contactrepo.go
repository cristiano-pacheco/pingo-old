// Package contactrepo contains the contact repository.
package contactrepo

import (
	"github.com/cristiano-pacheco/pingo/internal/domain/model/contactdm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
)

type ContactRepository interface {
	Create(user contactdm.Contact) error
	Update(user contactdm.Contact) error
	Delete(user contactdm.Contact) error
	FindByID(id identitydm.ID) (*contactdm.Contact, error)
	FindListByUserID(userID identitydm.ID) []*contactdm.Contact
}
