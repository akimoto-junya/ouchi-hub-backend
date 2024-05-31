package repository

import (
	"context"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/infra/db"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/repository"
	utilContext "github.com/akimoto-junya/ouchi-hub-backend/internal/util/context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FileType string

const (
	Text    = FileType("text")
	Image   = FileType("image")
	Audio   = FileType("audio")
	Video   = FileType("video")
	Unknown = FileType("unknown")
)

type content struct {
	pool    *pgxpool.Pool
	queries *db.Queries
}

func NewContent(pool *pgxpool.Pool) repository.Content {
	return &content{
		pool:    pool,
		queries: db.New(pool),
	}
}

func (c *content) Create(ctx context.Context, content *model.Content) error {
	if err := c.queries.CreateContent(ctx, db.CreateContentParams{
		ID:       content.ID.String(),
		ItemID:   content.ItemID.String(),
		Path:     content.Path,
		FileType: string(content.FileType),
	}); err != nil {
		return err
	}
	return nil
}

func (c *content) BatchCreate(ctx context.Context, contents []*model.Content) error {
	tx, ok := utilContext.GetTx(ctx)
	queries := func() *db.Queries {
		if ok {
			return c.queries.WithTx(tx)
		}
		return c.queries
	}()

	ids := make([]string, len(contents))
	itemIDs := make([]string, len(contents))
	paths := make([]string, len(contents))
	fileTypes := make([]string, len(contents))
	for i, content := range contents {
		ids[i] = content.ID.String()
		itemIDs[i] = content.ItemID.String()
		paths[i] = content.Path
		fileTypes[i] = string(content.FileType)
	}
	if err := queries.CreateContents(ctx, db.CreateContentsParams{
		Ids:       ids,
		ItemIds:   itemIDs,
		Paths:     paths,
		FileTypes: fileTypes,
	}); err != nil {
		return err
	}
	return nil
}

func (c *content) ListByItemID(ctx context.Context, itemID uuid.UUID) ([]*model.Content, error) {
	dbContents, err := c.queries.FindByItemID(ctx, itemID.String())
	if err != nil {
		return nil, err
	}
	contents := make([]*model.Content, len(dbContents))
	for i, dbContent := range dbContents {
		contents[i] = &model.Content{
			ID:       uuid.MustParse(dbContent.ID),
			ItemID:   uuid.MustParse(dbContent.ItemID),
			Path:     dbContent.Path,
			FileType: AsFileType(dbContent.FileType),
		}
	}
	return contents, nil
}

func AsFileType(filetype string) model.FileType {
	switch filetype {
	case string(Text):
		return model.Text
	case string(Image):
		return model.Image
	case string(Audio):
		return model.Audio
	case string(Video):
		return model.Video
	}
	return model.Unknown
}
