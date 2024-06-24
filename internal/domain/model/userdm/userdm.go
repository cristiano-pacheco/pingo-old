package userdm

import (
	"time"

	"github.com/cristiano-pacheco/pingo/internal/domain/model/errordm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
)

type User struct {
	ID                 identitydm.ID
	Name               Name
	Email              Email
	PasswordHash       []byte
	Status             Status
	ResetPasswordToken string
	CreatedAT          time.Time
	UpdatedAT          time.Time
}

func NewUser(name, email string, passwordHash []byte) (*User, error) {
	errs := errordm.New()
	n, err := NewName(name)
	errs.Add("name", err.Error())

	em, err := NewEmail(email)
	errs.Add("email", err.Error())

	status, err := NewStatus(STATUS_PENDING)
	errs.Add("status", err.Error())

	if errs.String() != "" {
		return nil, errs.Error()
	}

	return &User{
		ID:           *identitydm.New(),
		Name:         *n,
		Email:        *em,
		Status:       *status,
		PasswordHash: passwordHash,
		CreatedAT:    time.Now().UTC(),
	}, nil
}

func RestoreUser(
	id, name, email, passwordHash, status, resetPasswordToken string, createdAT, updatedAT time.Time,
) (*User, error) {
	errs := errordm.New()

	idVo, err := identitydm.Restore(id)
	errs.Add("id", err.Error())

	nameVo, err := NewName(name)
	errs.Add("name", err.Error())

	emailVo, err := NewEmail(email)
	errs.Add("email", err.Error())

	statusVo, err := NewStatus(status)
	errs.Add("status", err.Error())

	if errs.String() != "" {
		return nil, errs.Error()
	}

	user := &User{
		ID:           *idVo,
		Name:         *nameVo,
		Email:        *emailVo,
		PasswordHash: []byte(passwordHash),
		Status:       *statusVo,
		CreatedAT:    createdAT,
		UpdatedAT:    updatedAT,
	}

	return user, nil
}
