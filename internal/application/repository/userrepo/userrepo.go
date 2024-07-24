// Package userrepo contains the user repository contract interface.
package userrepo

import (
	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/userdm"
)

type UserRepository interface {
	Create(user userdm.User) error
	Update(user userdm.User) error
	Delete(user userdm.User) error
	UpdatePassword(user userdm.User) error
	UpdateResetPasswordToken(user userdm.User) error
	ActivateAccount(user userdm.User) error
	FindByID(id identitydm.ID) (*userdm.User, error)
	FindByEmail(email userdm.Email) (*userdm.User, error)
}
