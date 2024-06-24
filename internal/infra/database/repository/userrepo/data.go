package userrepo

import "time"

type UserDB struct {
	ID                 string
	Name               string
	Email              string
	PasswordHash       string
	Status             string
	ResetPasswordToken string
	CreatedAT          time.Time
	UpdatedAT          time.Time
}
