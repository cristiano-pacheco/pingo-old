package contactrepo

import (
	"time"
)

type ContactDB struct {
	ID          string
	UserID      string
	Name        string
	ContactType string
	ContactData string
	IsEnabled   bool
	CreatedAT   time.Time
	UpdatedAT   time.Time
}
