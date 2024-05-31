package model

import "github.com/google/uuid"

type Content struct {
	ID       uuid.UUID
	ItemID   uuid.UUID
	Path     string
	FileType FileType
}

type FileType string

const (
	Text    = FileType("text")
	Image   = FileType("image")
	Video   = FileType("video")
	Audio   = FileType("audio")
	Unknown = FileType("unknown")
)

func (t FileType) String() string {
	return string(t)
}
