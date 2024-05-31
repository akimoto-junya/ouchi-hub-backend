package repository

import (
	"context"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/infra/db"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/repository"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/infra/filesystem"
	utilContext "github.com/akimoto-junya/ouchi-hub-backend/internal/util/context"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/util/pointer"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type item struct {
	storage filesystem.OuchiHubStorage
	pool    *pgxpool.Pool
	queries *db.Queries
}

func NewItem(s filesystem.OuchiHubStorage, pool *pgxpool.Pool) repository.Item {
	return &item{
		storage: s,
		pool:    pool,
		queries: db.New(pool),
	}
}

func (i *item) GetDirectory(ctx context.Context, targetDirID uuid.UUID, depth int) (*model.Directory, error) {
	tx, ok := utilContext.GetTx(ctx)
	queries := func() *db.Queries {
		if ok {
			return i.queries.WithTx(tx)
		}
		return i.queries
	}()

	dbItems, err := queries.FindItemsWithDepth(ctx, db.FindItemsWithDepthParams{
		Depth:  depth,
		ItemID: targetDirID.String(),
	})
	if err != nil {
		return nil, err
	}
	directories := map[string]*model.Directory{}
	files := map[string]*model.File{}
	fileIDs := []string{}
	for _, dbItem := range dbItems {
		if _, ok := directories[dbItem.ParentID]; !ok && dbItem.Lft != 1 {
			directories[dbItem.ParentID] = &model.Directory{
				ID: uuid.MustParse(dbItem.ParentID),
				// Nameはここでは入れられない
				Directorires: []*model.Directory{},
				Files:        []*model.File{},
			}
		}
		// file
		if dbItem.IsFile {
			file := &model.File{
				ID:   uuid.MustParse(dbItem.ID),
				Name: dbItem.Name,
			}
			directories[dbItem.ParentID].Files = append(directories[dbItem.ParentID].Files, file)
			files[dbItem.ID] = file
			fileIDs = append(fileIDs, dbItem.ID)
			continue
		}
		// dir
		if _, ok := directories[dbItem.ID]; ok {
			directories[dbItem.ID].Name = dbItem.Name
		} else {
			dir := &model.Directory{
				ID:           uuid.MustParse(dbItem.ID),
				Name:         dbItem.Name,
				Directorires: []*model.Directory{},
				Files:        []*model.File{},
			}
			directories[dbItem.ID] = dir
		}
		if dbItem.Lft != 1 {
			directories[dbItem.ParentID].Directorires = append(directories[dbItem.ParentID].Directorires, directories[dbItem.ID])
		}
	}
	contents, err := queries.FindLatestContentsByItemIDs(ctx, fileIDs)
	if err != nil {
		return nil, err
	}
	for _, content := range contents {
		files[content.ItemID].Entity = &model.Content{
			ID:       uuid.MustParse(content.ID),
			ItemID:   uuid.MustParse(content.ItemID),
			Path:     content.Path,
			FileType: AsFileType(content.FileType),
		}
	}

	return directories[targetDirID.String()], nil
}

func (i *item) CreateRootDirectory(ctx context.Context, root *model.Directory, bucketID uuid.UUID) (retErr error) {
	if root == nil {
		return repository.ErrEmptyRootDir
	}
	tx, ok := utilContext.GetTx(ctx)
	queries := func() *db.Queries {
		if ok {
			return i.queries.WithTx(tx)
		}
		return i.queries
	}()

	// 入れ子集合モデルでのlft, rgtを生成しDB用のデータにする
	ranges := []*db.Item{}
	makeItemRanges(1, root, bucketID, &ranges, pointer.ToPtr(0))
	ids := make([]string, len(ranges))
	names := make([]string, len(ranges))
	bucketIDs := make([]string, len(ranges))
	isFiles := make([]bool, len(ranges))
	lfts := make([]int64, len(ranges))
	rgts := make([]int64, len(ranges))
	for i, dbItem := range ranges {
		ids[i] = dbItem.ID
		names[i] = dbItem.Name
		bucketIDs[i] = dbItem.BucketID
		isFiles[i] = dbItem.IsFile
		lfts[i] = dbItem.Lft
		rgts[i] = dbItem.Rgt
	}

	if err := queries.InsertItems(ctx, db.InsertItemsParams{
		Ids:       ids,
		Names:     names,
		BucketIds: bucketIDs,
		IsFiles:   isFiles,
		Lfts:      lfts,
		Rgts:      rgts,
	}); err != nil {
		return err
	}

	return nil
}

func (i *item) InsertItemTree(ctx context.Context, childRoot model.Item, bucketID, targetDirID uuid.UUID) (retErr error) {
	tx, ok := utilContext.GetTx(ctx)
	queries := func() *db.Queries {
		if ok {
			return i.queries.WithTx(tx)
		}
		return i.queries
	}()

	target, err := queries.FindItemByID(ctx, targetDirID.String())
	if err != nil {
		return err
	}

	ranges := []*db.Item{}
	count := 0
	makeItemRanges(target.Rgt+1, childRoot, bucketID, &ranges, &count)

	if err := queries.SpreadItemRange(ctx, db.SpreadItemRangeParams{
		Rgt: target.Rgt,
		Num: count,
	}); err != nil {
		return err
	}

	ids := make([]string, len(ranges))
	names := make([]string, len(ranges))
	bucketIDs := make([]string, len(ranges))
	lfts := make([]int64, len(ranges))
	rgts := make([]int64, len(ranges))
	for i, r := range ranges {
		ids[i] = r.ID
		names[i] = r.Name
		bucketIDs[i] = r.BucketID
		lfts[i] = r.Lft
		rgts[i] = r.Rgt
	}

	if err := queries.InsertItems(ctx, db.InsertItemsParams{
		Ids:       ids,
		Names:     names,
		BucketIds: bucketIDs,
		Lfts:      lfts,
		Rgts:      rgts,
	}); err != nil {
		return err
	}
	return nil
}

func makeItemRanges(lft int64, item model.Item, bucketID uuid.UUID, result *[]*db.Item, count *int) int64 {
	rgt := lft + 1
	if item.IsFile() {
		file := item.(*model.File)
		*result = append(*result, &db.Item{
			ID:       file.ID.String(),
			Name:     file.Name,
			BucketID: bucketID.String(),
			IsFile:   true,
			Lft:      lft,
			Rgt:      rgt,
		})
		*count++
		return rgt + 1
	}

	dir := item.(*model.Directory)
	rootLft := lft
	for _, d := range dir.Directorires {
		lft = rgt
		rgt = makeItemRanges(lft, d, bucketID, result, count)
	}
	for _, f := range dir.Files {
		lft = rgt
		rgt = makeItemRanges(lft, f, bucketID, result, count)
	}
	*result = append(*result, &db.Item{
		ID:       dir.ID.String(),
		Name:     dir.Name,
		BucketID: bucketID.String(),
		IsFile:   false,
		Lft:      rootLft,
		Rgt:      rgt,
	})
	*count++
	return rgt + 1
}

func (i *item) DeleteItemTree(ctx context.Context, bucketID, id uuid.UUID) (retErr error) {
	tx, ok := utilContext.GetTx(ctx)
	queries := func() *db.Queries {
		if ok {
			return i.queries.WithTx(tx)
		}
		return i.queries
	}()

	deleteItem, err := queries.FindItemByID(ctx, id.String())
	if err != nil {
		return err
	}

	if err := queries.DeleteItemTree(ctx, db.DeleteItemTreeParams{
		BucketID: bucketID.String(),
		ID:       id.String(),
	}); err != nil {
		return err
	}

	if err := queries.ShiftEmptyItemRange(ctx, db.ShiftEmptyItemRangeParams{
		Lft:      deleteItem.Lft,
		Rgt:      deleteItem.Rgt,
		BucketID: bucketID.String(),
	}); err != nil {
		return err
	}
	return nil
}

func (i *item) DeleteByBucketID(ctx context.Context, bucketID uuid.UUID) (retErr error) {
	tx, ok := utilContext.GetTx(ctx)
	queries := func() *db.Queries {
		if ok {
			return i.queries.WithTx(tx)
		}
		return i.queries
	}()
	if err := queries.DeleteByBucketID(ctx, bucketID.String()); err != nil {
		return err
	}
	return nil
}

func (i *item) LoadRootDirectory(ctx context.Context, bucketID uuid.UUID, relPath string) (*model.Directory, error) {
	return i.storage.FindWorkDir(bucketID, relPath)
}
