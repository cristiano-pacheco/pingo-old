package userrepo

import (
	"time"
)

type UserDB struct {
	ID                       string
	Name                     string
	Email                    string
	PasswordHash             []byte
	Status                   string
	ResetPasswordToken       []byte
	AccountConfirmationToken []byte
	CreatedAT                time.Time
	UpdatedAT                time.Time
}
