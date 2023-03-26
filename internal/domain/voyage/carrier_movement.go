package voyage

import (
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
