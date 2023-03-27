package location

import (
	"errors"

	"github.com/flowck/cargotracker_ddd_golang/internal/domain"
)

type Location struct {
	id                        domain.ID
	name                      string
	unitedNationsLocationCode UnitedNationsLocationCode
}

func New(
	ID domain.ID,
	name string,
	unitedNationsLocationCode UnitedNationsLocationCode,
) (*Location, error) {
	if ID.IsZero() {
		return nil, errors.New("id cannot be invalid")
	}

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if unitedNationsLocationCode.IsZero() {
		return nil, errors.New("unitedNationsLocationCode cannot be empty")
	}

	return &Location{
		id:                        ID,
		name:                      name,
		unitedNationsLocationCode: unitedNationsLocationCode,
	}, nil
}

func (l *Location) Id() domain.ID {
	return l.id
}

func (l *Location) UnitedNationsLocationCode() UnitedNationsLocationCode {
	return l.unitedNationsLocationCode
}

func (l *Location) Name() string {
	return l.name
}

func (l Location) IsEmpty() bool {
	return l.id.IsZero()
}
