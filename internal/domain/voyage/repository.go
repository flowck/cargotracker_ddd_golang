package voyage

import "context"

type Repository interface {
	Find(ctx context.Context, number Number) (*Voyage, error)
}
