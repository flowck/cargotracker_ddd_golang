package app

import (
	"context"
)

type CommandHandler[C any] interface {
	Execute(ctx context.Context, cmd C) error
}

type QueryHandler[Q any, R any] interface {
	Execute(ctx context.Context, q Q) (R, error)
}

type Commands struct {
}

type Queries struct {
}

type App struct {
	Commands Commands
	Queries  Queries
}
