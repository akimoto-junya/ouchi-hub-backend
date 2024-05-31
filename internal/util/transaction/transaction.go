package transaction

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Transact(ctx context.Context, pool *pgxpool.Pool) (pgx.Tx, func(error) error, error) {
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, nil, err
	}
	return tx, func(txErr error) error {
		if txErr != nil {
			if err := tx.Rollback(ctx); err != nil {
				return err
			}
			return txErr
		}
		if err := tx.Commit(ctx); err != nil {
			if err := tx.Rollback(ctx); err != nil {
				return err
			}
			return err
		}
		return nil
	}, nil
}
