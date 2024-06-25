// Package identitydm contains the ID domain model
package identitydm

import "github.com/google/uuid"

type ID struct {
	value string
}

func New() *ID {
	id := uuid.New()
	return &ID{value: id.String()}
}

func Restore(value string) (*ID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return nil, err
	}
	return &ID{value: id.String()}, nil
}

func (v *ID) String() string {
	return v.value
}
