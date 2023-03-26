package commands

import (
	"context"
)

type TransactableAdapters struct {
	EventPublisher EventPublisher
}

type EventPublisher interface {
	CargoHasArrived(ctx context.Context, event Cargo) error
	CargoWasHandled(ctx context.Context, event Handling) error
	CargoWasMisdirected(ctx context.Context, event Cargo) error
	ReceiveHandlingEventRegistration(ctx context.Context, event HandlingEventRegistrationAttempt) error
}

type TransactionProvider interface {
	Transact(ctx context.Context, f TransactFunc) error
}

type TransactFunc func(adapters TransactableAdapters) error
