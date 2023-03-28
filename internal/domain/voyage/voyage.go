package voyage

import (
	"errors"

	"github.com/flowck/cargotracker_ddd_golang/internal/domain"
)

//
// Voyage
//

type Voyage struct {
	id               domain.ID
	number           Number
	carrierMovements []CarrierMovement
}

func New(ID domain.ID, number Number, schedule Schedule) (*Voyage, error) {
	if ID.IsZero() {
		return nil, errors.New("id cannot be invalid")
	}

	if schedule.IsZero() {
		return nil, errors.New("schedule cannot be invalid")
	}

	if number.IsEmpty() {
		return nil, errors.New("number cannot be empty")
	}

	return &Voyage{
		id:               ID,
		number:           number,
		carrierMovements: schedule.CarrierMovements(),
	}, nil
}

func (v *Voyage) Id() domain.ID {
	return v.id
}

func (v *Voyage) Number() Number {
	return v.number
}

func (v *Voyage) CarrierMovements() []CarrierMovement {
	return v.carrierMovements
}
