package location

import "github.com/flowck/cargotracker_ddd_golang/internal/domain"

type Location struct {
	id                        domain.ID
	unitedNationsLocationCode UnitedNationsLocationCode
	name                      string
}

func (l Location) IsEmpty() bool {
	return l.id.IsZero()
}
