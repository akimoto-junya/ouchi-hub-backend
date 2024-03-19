package transaction

import (
	"context"

	utilContext "github.com/akimoto-junya/ouchi-hub-backend/internal/util/context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Transact(ctx context.Context, pool *pgxpool.Pool) (pgx.Tx, func() error, error) {
	var err error
	tx, ok := utilContext.GetTx(ctx)
	if !ok {
		tx, err = pool.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			return nil, nil, err
		}
		return tx, func() error {
			if err != nil {
				if err := tx.Rollback(ctx); err != nil {
					return err
				}
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
	return tx, func() error { return nil }, nil
}
