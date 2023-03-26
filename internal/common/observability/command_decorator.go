package observability

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/flowck/cargotracker_ddd_golang/internal/common/logs"
)

type commandDecorator[C any] struct {
	logger *logs.Logger
	base   CommandHandler[C]
}

func NewCommandDecorator[C any](
	base CommandHandler[C],
	logger *logs.Logger,
	tracer trace.Tracer,
) commandDecorator[C] {
	return commandDecorator[C]{
		logger: logger,
		base: commandMetricsDecorator[C]{
			base: commandTracingDecorator[C]{
				tracer: tracer,
				base: commandLoggingDecorator[C]{
					base:   base,
					logger: logger,
				},
			},
		},
	}
}

func (q commandDecorator[C]) Execute(ctx context.Context, cmd C) error {
	return q.base.Execute(ctx, cmd)
}
