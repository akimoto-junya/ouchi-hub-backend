package repository

import (
	"context"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/domain/model"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/domain/repository"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/infra/db"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/util/pointer"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/util/transaction"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type item struct {
	pool    *pgxpool.Pool
	queries *db.Queries
}

func NewRepository() repository.Item {
	return &item{
		// TODO:
	}
}

func (i *item) FindDirectChildren(ctx context.Context, targetDirID uuid.UUID) (ret []model.Item, retErr error) {
	tx, finishTx, err := transaction.Transact(ctx, i.pool)
	if err != nil {
		return nil, err
	}
	defer func() {
		retErr = finishTx()
	}()
	queries := i.queries.WithTx(tx)
	_, err = queries.FindDirectChildren(ctx, targetDirID.String())
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (i *item) CreateRootDirectory(ctx context.Context, root model.Directory, workID uuid.UUID) (retErr error) {
	tx, finishTx, err := transaction.Transact(ctx, i.pool)
	if err != nil {
		return err
	}
	defer func() {
		retErr = finishTx()
	}()
	queries := i.queries.WithTx(tx)

	// 入れ子集合モデルでのlft, rgtを生成しDB用のデータにする
	ranges := []*db.Item{}
	makeItemRanges(1, &root, workID, &ranges, pointer.ToPtr(0))
	ids := make([]string, len(ranges))
	workIDs := make([]string, len(ranges))
	lfts := make([]int64, len(ranges))
	rgts := make([]int64, len(ranges))
	for i, r := range ranges {
		ids[i] = r.ID
		workIDs[i] = r.WorkID
		lfts[i] = r.Lft
		rgts[i] = r.Rgt
	}

	if err := queries.InsertItems(ctx, db.InsertItemsParams{
		Ids:     ids,
		WorkIds: workIDs,
		Lfts:    lfts,
		Rgts:    rgts,
	}); err != nil {
		return err
	}

	return nil
}

func (i *item) DeleteItemTree(ctx context.Context, workID, id uuid.UUID) (retErr error) {
	tx, finishTx, err := transaction.Transact(ctx, i.pool)
	if err != nil {
		return err
	}
	defer func() {
		retErr = finishTx()
	}()
	queries := i.queries.WithTx(tx)

	deleteItem, err := queries.FindItemByID(ctx, id.String())
	if err != nil {
		return err
	}

	if err := queries.DeleteItemTree(ctx, db.DeleteItemTreeParams{
		WorkID: workID.String(),
		ID:     id.String(),
	}); err != nil {
		return err
	}

	if err := queries.ShiftEmptyItemRange(ctx, db.ShiftEmptyItemRangeParams{
		Lft:    deleteItem.Lft,
		Rgt:    deleteItem.Rgt,
		WorkID: workID.String(),
	}); err != nil {
		return err
	}
	return nil
}

func (i *item) InsertItemTree(ctx context.Context, childRoot model.Item, workID, targetDirID uuid.UUID) (retErr error) {
	tx, finishTx, err := transaction.Transact(ctx, i.pool)
	if err != nil {
		return err
	}
	defer func() {
		retErr = finishTx()
	}()
	queries := i.queries.WithTx(tx)

	target, err := queries.FindItemByID(ctx, targetDirID.String())
	if err != nil {
		return err
	}

	ranges := []*db.Item{}
	count := 0
	makeItemRanges(target.Rgt+1, childRoot, workID, &ranges, &count)

	if err := queries.SpreadItemRange(ctx, db.SpreadItemRangeParams{
		Rgt: target.Rgt,
		Num: count,
	}); err != nil {
		return err
	}

	ids := make([]string, len(ranges))
	names := make([]string, len(ranges))
	workIDs := make([]string, len(ranges))
	lfts := make([]int64, len(ranges))
	rgts := make([]int64, len(ranges))
	for i, r := range ranges {
		ids[i] = r.ID
		names[i] = r.Name
		workIDs[i] = r.WorkID
		lfts[i] = r.Lft
		rgts[i] = r.Rgt
	}

	if err := queries.InsertItems(ctx, db.InsertItemsParams{
		Ids:     ids,
		WorkIds: workIDs,
		Lfts:    lfts,
		Rgts:    rgts,
	}); err != nil {
		return err
	}
	return nil
}

func makeItemRanges(lft int64, item model.Item, workID uuid.UUID, result *[]*db.Item, count *int) int64 {
	rgt := lft + 1
	if item.IsFile() {
		file := item.(*model.File)
		*result = append(*result, &db.Item{
			ID:     file.ID.String(),
			Name:   file.Name,
			WorkID: workID.String(),
			Lft:    lft,
			Rgt:    rgt,
		})
		*count++
		return rgt + 1
	}

	dir := item.(*model.Directory)
	rootLft := lft
	for _, d := range dir.Directorires {
		lft = rgt
		rgt = makeItemRanges(lft, d, workID, result, count)
	}
	for _, f := range dir.Files {
		lft = rgt
		rgt = makeItemRanges(lft, f, workID, result, count)
	}
	*result = append(*result, &db.Item{
		ID:     dir.ID.String(),
		Name:   dir.Name,
		WorkID: workID.String(),
		Lft:    rootLft,
		Rgt:    rgt,
	})
	*count++
	return rgt + 1
}
