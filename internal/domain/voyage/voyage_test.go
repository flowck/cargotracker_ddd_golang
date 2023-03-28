package voyage_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/biter777/countries"
	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/flowck/cargotracker_ddd_golang/internal/domain"
	"github.com/flowck/cargotracker_ddd_golang/internal/domain/voyage"
)

func TestNewVoyage(t *testing.T) {
	gofakeit.Seed(0)

	testCases := []struct {
		name        string
		expectedErr string

		ID       domain.ID
		number   voyage.Number
		schedule voyage.Schedule
	}{
		{
			name:        "invalid_id",
			expectedErr: "id cannot be invalid",

			ID:       domain.ID{},
			number:   fixtureNumber(t),
			schedule: fixtureSchedule(t),
		},
		{
			name:        "empty_or_invalid_number",
			expectedErr: "number cannot be empty",

			ID:       domain.NewID(),
			number:   voyage.Number{},
			schedule: fixtureSchedule(t),
		},
		{
			name:        "invalid_schedule",
			expectedErr: "schedule cannot be invalid",

			ID:       domain.NewID(),
			number:   fixtureNumber(t),
			schedule: voyage.Schedule{},
		},
		{
			name:        "create_new_voyage",
			expectedErr: "",

			ID:       domain.NewID(),
			number:   fixtureNumber(t),
			schedule: fixtureSchedule(t),
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			v, err := voyage.New(
				tc.ID,
				tc.number,
				tc.schedule,
			)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.Nil(t, err)
			assert.Equal(t, tc.ID, v.Id())
			assert.Equal(t, tc.number, v.Number())
			assert.Equal(t, len(tc.schedule.CarrierMovements()), len(v.CarrierMovements()))

			for j, carrierMovement := range tc.schedule.CarrierMovements() {
				assert.Equal(t, carrierMovement.Id(), v.CarrierMovements()[j].Id())
			}
		})
	}
}

func fixtureSchedule(t *testing.T) voyage.Schedule {
	carrierMovements := make([]voyage.CarrierMovement, 10)
	for i := 0; i < 10; i++ {
		cm, err := voyage.NewCarrierMovement(
			domain.NewID(),
			*fixtureLocation(t),
			*fixtureLocation(t),
			time.Now(),
			time.Now().Add(time.Hour*20),
		)
		require.Nil(t, err)
		carrierMovements[i] = *cm
	}

	s, err := voyage.NewSchedule(carrierMovements)
	require.Nil(t, err)

	return s
}

func fixtureNumber(t *testing.T) voyage.Number {
	n, err := voyage.NewNumber(fmt.Sprintf(
		"%s-%d",
		countries.ByName(gofakeit.Country()),
		gofakeit.Number(100, 200),
	))

	require.Nil(t, err)
	return n
}
