package filesystem

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/model"
)

func DetectFileType(filePath string) model.FileType {
	f, err := os.Open(filePath)
	defer func() {
		_ = f.Close()
	}()
	if err != nil && errors.Is(err, io.EOF) {
		return detectFileTypeByExt(filePath)
	}
	head := make([]byte, 512)
	if _, err := f.Read(head); err != nil {
		return detectFileTypeByExt(filePath)
	}
	mimeType := http.DetectContentType(head)
	if strings.HasPrefix(mimeType, "text") {
		return model.Text
	} else if strings.HasPrefix(mimeType, "image") {
		return model.Image
	} else if strings.HasPrefix(mimeType, "audio") {
		return model.Audio
	} else if strings.HasPrefix(mimeType, "video") {
		return model.Video
	}
	return detectFileTypeByExt(filePath)
}

var exts = map[string]model.FileType{
	// text
	".txt":  model.Text,
	".text": model.Text,
	// image
	".jpg":  model.Image,
	".jpeg": model.Image,
	".png":  model.Image,
	".gif":  model.Image,
	".svg":  model.Image,
	".webp": model.Image,
	".avif": model.Image,
	// audio
	".mp3":  model.Audio,
	".wav":  model.Audio,
	".aac":  model.Audio,
	".flac": model.Audio,
	".oga":  model.Audio,
	// video
	".mp4":  model.Video,
	".webm": model.Video,
	".ogv":  model.Video,
}

func detectFileTypeByExt(filePath string) model.FileType {
	ext := strings.ToLower(filepath.Ext(filePath))
	if t, ok := exts[ext]; ok {
		return t
	}
	return model.Unknown
}
