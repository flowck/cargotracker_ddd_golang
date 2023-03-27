package voyage_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flowck/cargotracker_ddd_golang/internal/domain"
	"github.com/flowck/cargotracker_ddd_golang/internal/domain/location"
	"github.com/flowck/cargotracker_ddd_golang/internal/domain/voyage"
)

func TestNewCarrierMovement(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		ID                domain.ID
		arrivalLocation   location.Location
		departureLocation location.Location
		arrivalTime       time.Time
		departureTime     time.Time
	}{
		{
			name:        "error_invalid_id",
			expectedErr: "id cannot be invalid",

			ID:                domain.ID{},
			arrivalLocation:   *fixtureLocation(t),
			departureLocation: *fixtureLocation(t),
			arrivalTime:       time.Now(),
			departureTime:     time.Now().Add(time.Hour * 20),
		},
		{
			name:        "error_invalid_arrival_location",
			expectedErr: "arrivalLocation cannot be invalid",

			ID:                domain.NewID(),
			arrivalLocation:   location.Location{},
			departureLocation: *fixtureLocation(t),
			arrivalTime:       time.Now(),
			departureTime:     time.Now().Add(time.Hour * 20),
		},
		{
			name:        "error_invalid_departure_location",
			expectedErr: "departureLocation cannot be invalid",

			ID:                domain.NewID(),
			arrivalLocation:   *fixtureLocation(t),
			departureLocation: location.Location{},
			arrivalTime:       time.Now(),
			departureTime:     time.Now().Add(time.Hour * 20),
		},
		{
			name:        "error_arrival_is_in_the_future",
			expectedErr: "arrivalTime cannot be after departure time",

			ID:                domain.NewID(),
			arrivalLocation:   *fixtureLocation(t),
			departureLocation: *fixtureLocation(t),
			arrivalTime:       time.Now().Add(time.Hour * 20),
			departureTime:     time.Now(),
		},
		{
			name:        "create_new_carrier_movement",
			expectedErr: "",

			ID:                domain.NewID(),
			arrivalLocation:   *fixtureLocation(t),
			departureLocation: *fixtureLocation(t),
			arrivalTime:       time.Now(),
			departureTime:     time.Now().Add(time.Hour * 20),
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			carrierMovement, err := voyage.NewCarrierMovement(
				tc.ID,
				tc.arrivalLocation,
				tc.departureLocation,
				tc.arrivalTime,
				tc.departureTime,
			)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.Nil(t, err)
			assert.Equal(t, carrierMovement.Id(), tc.ID)
			assert.Equal(t, carrierMovement.ArrivalTime(), tc.arrivalTime)
			assert.Equal(t, carrierMovement.DepartureTime(), tc.departureTime)
			assert.Equal(t, carrierMovement.ArrivalLocation(), tc.arrivalLocation)
			assert.Equal(t, carrierMovement.DepartureLocation(), tc.departureLocation)
		})
	}
}

func fixtureLocation(t *testing.T) *location.Location {
	code, err := location.NewUnitedNationsLocationCodeFromCountry(gofakeit.Country())
	require.Nil(t, err)

	l, err := location.New(
		domain.NewID(),
		gofakeit.Country(),
		code,
	)
	require.Nil(t, err)
	return l
}
