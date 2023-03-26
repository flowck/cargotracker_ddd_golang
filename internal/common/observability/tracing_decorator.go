package observability

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

type commandTracingDecorator[C any] struct {
	base   CommandHandler[C]
	tracer trace.Tracer
}

type queryTracingDecorator[Q any, R any] struct {
	base   QueryHandler[Q, R]
	tracer trace.Tracer
}

func (c commandTracingDecorator[C]) Execute(ctx context.Context, cmd C) error {
	ctx, span := c.tracer.Start(ctx, handlerName(c.base))
	defer span.End()

	return c.base.Execute(ctx, cmd)
}

func (c queryTracingDecorator[Q, R]) Execute(ctx context.Context, q Q) (R, error) {
	ctx, span := c.tracer.Start(ctx, handlerName(c.base))
	defer span.End()

	return c.base.Execute(ctx, q)
}
