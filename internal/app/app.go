package app

import (
	"context"

	"github.com/flowck/cargotracker_ddd_golang/internal/app/commands"
)

type CommandHandler[C any] interface {
	Execute(ctx context.Context, cmd C) error
}

type QueryHandler[Q any, R any] interface {
	Execute(ctx context.Context, q Q) (R, error)
}

type Commands struct {
	// Booking commands
	BookNewCargo       CommandHandler[commands.BookNewCargo]
	AssignCargoToRoute CommandHandler[any]
	ChangeDestination  CommandHandler[any]

	// Cargo inspection commands
	InspectCargo CommandHandler[any]
}

type Queries struct {
	RequestPossibleRoutesForCargo QueryHandler[any, any]
}

type App struct {
	Commands Commands
	Queries  Queries
}
