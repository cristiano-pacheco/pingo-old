package userrepo

import (
	"database/sql"
	"time"
)

type UserDB struct {
	ID                 string
	Name               string
	Email              string
	PasswordHash       string
	Status             string
	ResetPasswordToken sql.NullString
	CreatedAT          time.Time
	UpdatedAT          time.Time
}
