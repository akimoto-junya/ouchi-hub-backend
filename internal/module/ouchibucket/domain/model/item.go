package model

import "github.com/google/uuid"

type Item interface {
	IsFile() bool
	IsDir() bool
}

type File struct {
	ID     uuid.UUID
	Name   string
	Entity *Content
}

func (f *File) IsFile() bool { return true }
func (f *File) IsDir() bool  { return false }

func NewFile(name string) *File {
	return &File{
		ID:     uuid.New(),
		Name:   name,
		Entity: nil,
	}
}

type Directory struct {
	ID           uuid.UUID
	Name         string
	Directorires []*Directory
	Files        []*File
}

func (d *Directory) IsFile() bool { return false }
func (d *Directory) IsDir() bool  { return true }

func NewDirectory(name string) *Directory {
	return &Directory{
		ID:           uuid.New(),
		Name:         name,
		Directorires: []*Directory{},
		Files:        []*File{},
	}
}

func (d *Directory) ListContentsRecursive() []*Content {
	return listContentsRecursive(d)
}

func listContentsRecursive(dir *Directory) []*Content {
	contents := []*Content{}
	for _, f := range dir.Files {
		contents = append(contents, f.Entity)
	}
	for _, d := range dir.Directorires {
		contents = append(contents, listContentsRecursive(d)...)
	}
	return contents
}
