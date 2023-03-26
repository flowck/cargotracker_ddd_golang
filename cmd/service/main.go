package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/flowck/cargotracker_ddd_golang/internal/ports/rpc"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"

	"github.com/flowck/cargotracker_ddd_golang/internal/app"
	"github.com/flowck/cargotracker_ddd_golang/internal/app/commands"
	"github.com/flowck/cargotracker_ddd_golang/internal/common/logs"
	"github.com/flowck/cargotracker_ddd_golang/internal/common/observability"
	"github.com/flowck/cargotracker_ddd_golang/internal/common/psql"
	"github.com/flowck/cargotracker_ddd_golang/internal/ports/http"
)

const (
	ServiceName = "cargotracker"
	Version     = ""
)

type Config struct {
	Port              int16  `envconfig:"PORT"`
	GrpcPort          int16  `envconfig:"GRPC_PORT"`
	DebugMode         string `envconfig:"FLAG_DEBUG_MODE"`
	AllowedCorsOrigin string `envconfig:"ALLOWED_CORS_ORIGIN"`
	DatabaseUrl       string `envconfig:"DATABASE_URL"`
}

func main() {
	cfg, err := getConfig()
	if err != nil {
		panic(err)
	}

	logger := logs.New(cfg.DebugMode == "enabled")
	logger.WithFields(logs.Fields{"ServiceName": ServiceName, "version": Version}).Info()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	db, err := psql.Connect(cfg.DatabaseUrl)
	if err != nil {
		logger.Fatal(err)
	}

	err = psql.ApplyMigrations(db, "misc/sql/migrations")
	if err != nil {
		logger.Fatal(err)
	}

	// txProvider := transaction.NewSQLProvider(db, logger)

	//
	// OpenTelemetry Tracing
	//
	traceProvider, err := newOtelTraceProvider(nil)
	if err != nil {
		logger.Fatalf("unable to create trace provider: %v", err)
	}
	defer func() {
		logger.Infof("trace provider has been shutdown: %v", traceProvider.Shutdown(ctx))
	}()
	otel.SetTracerProvider(traceProvider)
	tracer := traceProvider.Tracer(ServiceName)

	application := &app.App{
		Commands: app.Commands{
			BookNewCargo: observability.NewCommandDecorator[commands.BookNewCargo](commands.NewBookNewCargo(), logger, tracer),
		},
		Queries: app.Queries{},
	}

	httpPort := http.NewPort(ctx, cfg.Port, strings.Split(cfg.AllowedCorsOrigin, ";"), application, logger)
	go func() { httpPort.Start() }()

	rpcPort := rpc.NewPort(logger)
	go func() {
		if err = rpcPort.Start(cfg.GrpcPort); err != nil {
			logger.Fatal(err)
		}
	}()

	<-done
	logger.Info("Preparing to shutdown gracefully")

	ctxTerm, cancelCtxTerm := context.WithTimeout(ctx, time.Second*15)
	defer cancelCtxTerm()

	httpPort.Stop(ctxTerm)
	rpcPort.Stop()

	logger.Info("The service has been terminated")
}

func getConfig() (*Config, error) {
	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

/*func newOtelExporter() sdktrace.SpanExporter {
	return nil
}*/

func newOtelTraceProvider(exp sdktrace.SpanExporter) (*sdktrace.TracerProvider, error) {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(ServiceName),
		),
	)

	if err != nil {
		return nil, err
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(r),
	), nil
}
