package voyage

import "errors"

//
// A valueObject to represent the voyage's schedule
//

type Schedule struct {
	carrierMovements []CarrierMovement
}

func NewSchedule(carrierMovements []CarrierMovement) (Schedule, error) {
	if carrierMovements == nil {
		return Schedule{}, errors.New("carrierMovements cannot be nil")
	}

	if len(carrierMovements) == 0 {
		return Schedule{}, errors.New("carrierMovements cannot be empty")
	}

	return Schedule{
		carrierMovements: carrierMovements,
	}, nil
}

func (s Schedule) CarrierMovements() []CarrierMovement {
	return s.carrierMovements
}

func (s Schedule) IsZero() bool {
	return s.carrierMovements == nil
}
