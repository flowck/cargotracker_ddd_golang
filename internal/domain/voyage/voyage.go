package voyage

import "github.com/flowck/cargotracker_ddd_golang/internal/domain"

type Voyage struct {
	id               domain.ID
	voyageNumber     string
	carrierMovements []CarrierMovement
}
