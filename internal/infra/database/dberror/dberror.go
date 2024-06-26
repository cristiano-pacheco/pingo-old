// Package dberror contains the errors related to the database.
package dberror

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
)

const ErrUniqueViolationCode = "23505"
