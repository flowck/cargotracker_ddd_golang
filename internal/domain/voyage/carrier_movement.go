package voyage

import (
	"errors"
	"time"

	"github.com/flowck/cargotracker_ddd_golang/internal/domain"
	"github.com/flowck/cargotracker_ddd_golang/internal/domain/location"
)

type CarrierMovement struct {
	id                domain.ID
	arrivalLocation   location.Location
	departureLocation location.Location
	arrivalTime       time.Time
	departureTime     time.Time
}

func NewCarrierMovement(
	ID domain.ID,
	arrivalLocation location.Location,
	departureLocation location.Location,
	arrivalTime time.Time,
	departureTime time.Time,
) (*CarrierMovement, error) {
	if ID.IsZero() {
		return nil, errors.New("id cannot be invalid")
	}

	if arrivalLocation.IsEmpty() {
		return nil, errors.New("arrivalLocation cannot be invalid")
	}

	if departureLocation.IsEmpty() {
		return nil, errors.New("departureLocation cannot be invalid")
	}

	if arrivalTime.After(departureTime) {
		return nil, errors.New("arrivalTime cannot be after departure time")
	}

	return &CarrierMovement{
		id:                ID,
		arrivalLocation:   arrivalLocation,
		departureLocation: departureLocation,
		arrivalTime:       arrivalTime,
		departureTime:     departureTime,
	}, nil
}

func (c *CarrierMovement) Id() domain.ID {
	return c.id
}

func (c *CarrierMovement) ArrivalLocation() location.Location {
	return c.arrivalLocation
}

func (c *CarrierMovement) DepartureLocation() location.Location {
	return c.departureLocation
}

func (c *CarrierMovement) ArrivalTime() time.Time {
	return c.arrivalTime
}

func (c *CarrierMovement) DepartureTime() time.Time {
	return c.departureTime
}

func (c CarrierMovement) IsEmpty() bool {
	return c.id.IsZero()
}
