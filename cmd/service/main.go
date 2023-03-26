package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/flowck/cargotracker_ddd_golang/internal/app"
	"github.com/flowck/cargotracker_ddd_golang/internal/common/logs"
	"github.com/flowck/cargotracker_ddd_golang/internal/common/psql"
	"github.com/flowck/cargotracker_ddd_golang/internal/ports/http"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

type Config struct {
	Port              int16  `envconfig:"PORT"`
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
	logger.Info("cargotracker")

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

	application := &app.App{
		Commands: app.Commands{},
		Queries:  app.Queries{},
	}

	httpPort := http.NewPort(ctx, cfg.Port, strings.Split(cfg.AllowedCorsOrigin, ";"), application, logger)
	httpPort.Start()

	<-done
	logger.Info("Preparing to shutdown gracefully")

	ctxTerm, cancelCtxTerm := context.WithTimeout(ctx, time.Second*15)
	defer cancelCtxTerm()

	httpPort.Stop(ctxTerm)

	logger.Info("The service has been terminated")
}

func getConfig() (*Config, error) {
	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
