package cargo

import (
	"time"

	"github.com/flowck/cargotracker_ddd_golang/internal/domain"
	"github.com/flowck/cargotracker_ddd_golang/internal/domain/location"
)

type Voyage struct{}

type Leg struct {
	id             domain.ID
	voyage         Voyage
	loadLocation   location.Location
	unloadLocation location.Location
	loadTime       time.Time
	unloadTime     time.Time
}

func NewLeg(
	ID domain.ID,
	voyage Voyage,
	loadLocation location.Location,
	unloadLocation location.Location,
	loadTime time.Time,
	unloadTime time.Time,
) (*Leg, error) {
	return &Leg{
		id:             ID,
		voyage:         voyage,
		loadLocation:   loadLocation,
		unloadLocation: unloadLocation,
		loadTime:       loadTime,
		unloadTime:     unloadTime,
	}, nil
}

func (l *Leg) ID() domain.ID {
	return l.id
}

func (l *Leg) Voyage() Voyage {
	return l.voyage
}

func (l *Leg) LoadLocation() location.Location {
	return l.loadLocation
}

func (l *Leg) UnloadLocation() location.Location {
	return l.unloadLocation
}

func (l *Leg) LoadTime() time.Time {
	return l.loadTime
}

func (l *Leg) UnloadTime() time.Time {
	return l.unloadTime
}

func (l *Leg) IsEmpty() bool {
	return l.id.IsZero()
}
