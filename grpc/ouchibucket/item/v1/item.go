package itemv1

import (
	contentv1 "github.com/akimoto-junya/ouchi-hub-backend/grpc/ouchibucket/content/v1"
	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
)

func ToDirectoryMessage(d *model.Directory) *Directory {
	return &Directory{
		Id:   d.ID.String(),
		Name: d.Name,
		Directories: func() []*Directory {
			dirs := make([]*Directory, len(d.Directorires))
			for i, dir := range d.Directorires {
				dirs[i] = ToDirectoryMessage(dir)
			}
			return dirs
		}(),
		Files: func() []*File {
			files := make([]*File, len(d.Files))
			for i, file := range d.Files {
				files[i] = ToFileMessage(file)
			}
			return files
		}(),
	}
}

func ToFileMessage(f *model.File) *File {
	return &File{
		Id:   f.ID.String(),
		Name: f.Name,
		FileType: func() contentv1.FileType {
			switch f.Entity.FileType.String() {
			case "text":
				return contentv1.FileType_FILE_TYPE_TEXT
			case "image":
				return contentv1.FileType_FILE_TYPE_IMAGE
			case "audio":
				return contentv1.FileType_FILE_TYPE_AUDIO
			case "video":
				return contentv1.FileType_FILE_TYPE_VIDEO
			}
			return contentv1.FileType_FILE_TYPE_UNSPECIFIED
		}(),
		Url: f.Entity.Path,
	}
}
