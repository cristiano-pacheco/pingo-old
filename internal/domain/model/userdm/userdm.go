// Package userdm contains the user domain model.
package userdm

import (
	"time"

	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
)

type User struct {
	ID                       identitydm.ID
	Name                     Name
	Email                    Email
	PasswordHash             []byte
	Status                   Status
	ResetPasswordToken       []byte
	AccountConfirmationToken []byte
	CreatedAT                time.Time
	UpdatedAT                time.Time
}

func NewUser(name, email string, passwordHash, accountConfToken []byte) (*User, error) {
	nameVo, err := NewName(name)
	if err != nil {
		return nil, err
	}

	emailVo, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	statusVo, err := NewStatus(StatusPending)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:                       *identitydm.New(),
		Name:                     *nameVo,
		Email:                    *emailVo,
		Status:                   *statusVo,
		PasswordHash:             passwordHash,
		AccountConfirmationToken: accountConfToken,
		CreatedAT:                time.Now().UTC(),
	}, nil
}

func RestoreUser(
	id, name, email, status string,
	passwordHash, accountConfToken, resetPasswordToken []byte,
	createdAT, updatedAT time.Time,
) (*User, error) {
	idVo, err := identitydm.Restore(id)
	if err != nil {
		return nil, err
	}

	nameVo, err := NewName(name)
	if err != nil {
		return nil, err
	}

	emailVo, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	statusVo, err := NewStatus(StatusPending)
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:                       *idVo,
		Name:                     *nameVo,
		Email:                    *emailVo,
		PasswordHash:             passwordHash,
		AccountConfirmationToken: accountConfToken,
		ResetPasswordToken:       resetPasswordToken,
		Status:                   *statusVo,
		CreatedAT:                createdAT,
		UpdatedAT:                updatedAT,
	}

	return user, nil
}

func (m *User) IsEnabled() bool {
	return m.Status.String() == StatusConfirmed
}
