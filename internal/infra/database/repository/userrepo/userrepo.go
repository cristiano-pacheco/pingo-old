// Package userrepo contains the user repository stuff.
package userrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/userdm"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user userdm.User) error {
	query := `INSERT INTO users 
	(id, name, email, password_hash, account_confirmation_token, status, created_at, updated_at) 
	values ($1, $2, $3, $4, $5, $6, now(), now())`

	args := []any{
		user.ID.String(),
		user.Name.String(),
		user.Email.String(),
		user.PasswordHash,
		user.AccountConfirmationToken,
		user.Status.String(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(user userdm.User) error {
	query := `UPDATE users set name = $1, email = $2, status = $3 where id = $4`

	args := []any{
		user.Name.String(),
		user.Email.String(),
		user.Status.String(),
		user.ID.String(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdatePassword(user userdm.User) error {
	query := `UPDATE users set password_hash = $1 where id = $2`

	args := []any{
		user.PasswordHash,
		user.ID.String(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateResetPasswordToken(user userdm.User) error {
	query := `UPDATE users set reset_password_token = $1 where id = $2`

	args := []any{
		user.ResetPasswordToken,
		user.ID.String(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) ActivateAccount(user userdm.User) error {
	query := `UPDATE users set account_confirmation_token = null, status = $1 where id = $2`

	args := []any{
		userdm.StatusConfirmed,
		user.ID.String(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(user userdm.User) error {
	query := `DELETE from users where id = $1`

	args := []any{
		user.ID.String(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindByID(id identitydm.ID) (*userdm.User, error) {
	query := `
		select id, name, email, password_hash, status, account_confirmation_token, reset_password_token, created_at, updated_at
		from users where id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var userdb UserDB

	err := r.db.QueryRowContext(ctx, query, id.String()).Scan(
		&userdb.ID,
		&userdb.Name,
		&userdb.Email,
		&userdb.PasswordHash,
		&userdb.Status,
		&userdb.AccountConfirmationToken,
		&userdb.ResetPasswordToken,
		&userdb.CreatedAT,
		&userdb.UpdatedAT,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, fmt.Errorf("the user with ID: %s is not found", id.String())
		default:
			return nil, err
		}
	}

	user, err := mapUserDBToUser(&userdb)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByEmail(email userdm.Email) (*userdm.User, error) {
	query := `
		select id, name, email, password_hash, status, account_confirmation_token reset_password_token, created_at, updated_at
		from users where email = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var userdb UserDB

	err := r.db.QueryRowContext(ctx, query, email.String()).Scan(
		&userdb.ID,
		&userdb.Name,
		&userdb.Email,
		&userdb.PasswordHash,
		&userdb.Status,
		&userdb.AccountConfirmationToken,
		&userdb.ResetPasswordToken,
		&userdb.CreatedAT,
		&userdb.UpdatedAT,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, fmt.Errorf("the user with Email: %s is not found", email.String())
		default:
			return nil, err
		}
	}

	user, err := mapUserDBToUser(&userdb)
	if err != nil {
		return nil, err
	}

	return user, nil
}
