package context

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type TxKey struct{}

func GetTx(ctx context.Context) (pgx.Tx, bool) {
	if tx, ok := ctx.Value(TxKey{}).(pgx.Tx); ok {
		return tx, true
	}
	return nil, false
}

func SetTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey{}, tx)
}
