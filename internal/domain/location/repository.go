package location

import "context"

type Repository interface {
	Find(ctx context.Context, unitedNationsLocationCode UnitedNationsLocationCode) (*Location, error)
	FindAll(ctx context.Context) ([]*Location, error)
}
