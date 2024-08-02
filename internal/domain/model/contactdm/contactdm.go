// Package contactdm contains contact domain model.
package contactdm

import (
	"time"

	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
)

type Contact struct {
	ID     identitydm.ID
	UserID identitydm.ID
	Name   Name
	ContactData
	IsEnabled bool
	CreatedAT time.Time
	UpdatedAT time.Time
}

func New(userID, name, contactType, contactValue string, isEnabled bool) (*Contact, error) {
	idVo := identitydm.New()

	userIDVo, err := identitydm.Restore(userID)
	if err != nil {
		return nil, err
	}

	nameVo, err := NewName(name)
	if err != nil {
		return nil, err
	}

	contactData, err := NewContactData(contactType, contactValue)
	if err != nil {
		return nil, err
	}

	c := &Contact{
		ID:          *idVo,
		UserID:      *userIDVo,
		Name:        *nameVo,
		ContactData: *contactData,
		IsEnabled:   isEnabled,
		CreatedAT:   time.Now().UTC(),
	}

	return c, nil
}

func Restore(
	id, userID, name, contactType, contactValue string,
	isEnabled bool, createdAT, updatedAT time.Time,
) (*Contact, error) {
	idVo, err := identitydm.Restore(id)
	if err != nil {
		return nil, err
	}

	userIDVo, err := identitydm.Restore(userID)
	if err != nil {
		return nil, err
	}

	nameVo, err := NewName(name)
	if err != nil {
		return nil, err
	}

	contactData, err := NewContactData(contactType, contactValue)
	if err != nil {
		return nil, err
	}

	c := &Contact{
		ID:          *idVo,
		UserID:      *userIDVo,
		Name:        *nameVo,
		ContactData: *contactData,
		IsEnabled:   isEnabled,
		CreatedAT:   createdAT,
		UpdatedAT:   updatedAT,
	}

	return c, nil
}
