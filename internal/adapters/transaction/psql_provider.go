package transaction

import (
	"context"
	"database/sql"

	"github.com/flowck/cargotracker_ddd_golang/internal/app/commands"
	"github.com/flowck/cargotracker_ddd_golang/internal/common/logs"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type SQLProvider struct {
	db     boil.ContextBeginner
	logger *logs.Logger
}

func NewSQLProvider(db boil.ContextBeginner, logger *logs.Logger) SQLProvider {
	return SQLProvider{
		db:     db,
		logger: logger,
	}
}

func (p SQLProvider) Transact(ctx context.Context, f commands.TransactFunc) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	adapters := commands.TransactableAdapters{}

	err = f(adapters)
	if err != nil {
		mustRollbackTx(p.logger, tx, err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func mustRollbackTx(logger *logs.Logger, tx *sql.Tx, err error) {
	if rollbackErr := tx.Rollback(); rollbackErr != nil {
		logger.WithError(err).WithField("rollback_err", rollbackErr).Error("Rollback error")
	}
}
