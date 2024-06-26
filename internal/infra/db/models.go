// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Bucket struct {
	ID        string
	Name      string
	StorageID string
}

type Category struct {
	ID   string
	Name string
}

type Content struct {
	ID        string
	ItemID    string
	Path      string
	FileType  string
	CreatedAt pgtype.Timestamp
}

type Item struct {
	ID       string
	Name     string
	BucketID string
	IsFile   bool
	Lft      int64
	Rgt      int64
}

type Maker struct {
	ID   string
	Name string
}

type Storage struct {
	ID      string
	OwnerID string
}

type User struct {
	ID string
}

type Work struct {
	ID         string
	Title      string
	CategoryID string
	MakerID    string
	BucketID   string
}
