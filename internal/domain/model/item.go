package model

import "github.com/google/uuid"

type Item interface {
	IsFile() bool
	IsDir() bool
}

type File struct {
	ID   uuid.UUID
	Name string
}

func (f *File) IsFile() bool { return true }
func (f *File) IsDir() bool  { return false }

type Directory struct {
	ID           uuid.UUID
	Name         string
	Directorires []*Directory
	Files        []*File
}

func (d *Directory) IsFile() bool { return false }
func (d *Directory) IsDir() bool  { return true }
