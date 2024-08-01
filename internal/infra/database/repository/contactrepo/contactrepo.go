// Package contactrepo contains the contact repository.
package contactrepo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/cristiano-pacheco/pingo/internal/domain/model/contactdm"
	"github.com/cristiano-pacheco/pingo/internal/domain/model/identitydm"
)

type ContactRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) Create(contact contactdm.Contact) error {
	query := `INSERT INTO contact 
	(id, user_id, name, contact_type, contact_data, is_enabled, created_at, updated_at) 
	values ($1, $2, $3, $4, $5, $6, now(), now())`

	args := []any{
		contact.ID.String(),
		contact.UserID.String(),
		contact.Name.String(),
		contact.ContactData.ContactType(),
		contact.ContactData.ContactValue(),
		contact.IsEnabled,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *ContactRepository) Update(contact contactdm.Contact) error {
	query := `UPDATE contact set name = $1, contact_type = $2, contact_data = $3, is_enabled = $4, updated_at = now() where id = $5 and user_id = $6`

	args := []any{
		contact.Name.String(),
		contact.ContactData.ContactType(),
		contact.ContactData.ContactValue(),
		contact.IsEnabled,
		contact.ID.String(),
		contact.UserID.String(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *ContactRepository) Delete(contact contactdm.Contact) error {
	query := `DELETE from contact where id = $1 and user_id = $2`

	args := []any{
		contact.ID.String(),
		contact.UserID.String(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *ContactRepository) FindByIDAndUserID(id, userID identitydm.ID) (*contactdm.Contact, error) {
	query := `
		select id, user_id, name, contact_type, contact_data, is_enabled, created_at, updated_at
		from contact where id = $1 and user_id = $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var contactdb ContactDB

	err := r.db.QueryRowContext(ctx, query, id.String()).Scan(
		&contactdb.ID,
		&contactdb.UserID,
		&contactdb.Name,
		&contactdb.ContactType,
		&contactdb.ContactData,
		&contactdb.IsEnabled,
		&contactdb.CreatedAT,
		&contactdb.UpdatedAT,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, fmt.Errorf(
				"the contact with ID: %s and user_id: %s is not found",
				id.String(),
				userID.String(),
			)
		default:
			return nil, err
		}
	}

	user, err := mapContactDBToContact(&contactdb)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *ContactRepository) FindListByUserID(userID identitydm.ID) ([]*contactdm.Contact, error) {
	query := `
		select id, user_id, name, contact_type, contact_data, is_enabled, created_at, updated_at
		from contact where user_id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, userID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []*contactdm.Contact

	for rows.Next() {
		var contactdb ContactDB
		err := rows.Scan(
			&contactdb.ID,
			&contactdb.UserID,
			&contactdb.Name,
			&contactdb.ContactType,
			&contactdb.ContactData,
			&contactdb.IsEnabled,
			&contactdb.CreatedAT,
			&contactdb.UpdatedAT,
		)
		if err != nil {
			return nil, err
		}

		contact, err := mapContactDBToContact(&contactdb)
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return contacts, nil
}
