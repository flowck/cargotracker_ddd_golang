package cargo

import (
	"github.com/flowck/cargotracker_ddd_golang/internal/domain"
	"github.com/flowck/cargotracker_ddd_golang/internal/domain/location"
)

type HandlingHistory struct{}

type Itinerary struct {
	legs []Leg
}

func (i *Itinerary) Legs() []Leg {
	return i.legs
}

type RouteSpecification struct{}

type Cargo struct {
	id                 domain.ID
	trackingID         domain.ID
	itineraryLegs      []Leg
	origin             location.Location
	delivery           Delivery
	routeSpecification RouteSpecification
}

func (c *Cargo) Id() domain.ID {
	return c.id
}

func (c *Cargo) TrackingID() domain.ID {
	return c.trackingID
}

func (c *Cargo) Origin() location.Location {
	return c.origin
}

func (c *Cargo) Delivery() Delivery {
	return c.delivery
}

func (c *Cargo) RouteSpecification() RouteSpecification {
	return c.routeSpecification
}

func (c *Cargo) assignToRoute(itinerary Itinerary) error {
	return nil
}

func (c Cargo) deriveDeliveryProgress(handlingHistory HandlingHistory) {
}

func New(
	ID domain.ID,
	trackingID domain.ID,
	routeSpecification RouteSpecification,
	itinerary Itinerary,
) (*Cargo, error) {
	return &Cargo{
		id:                 ID,
		trackingID:         trackingID,
		itineraryLegs:      itinerary.legs,
		origin:             location.Location{},
		delivery:           Delivery{},
		routeSpecification: routeSpecification,
	}, nil
}
