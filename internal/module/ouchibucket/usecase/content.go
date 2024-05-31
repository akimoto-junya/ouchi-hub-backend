package usecase

import "github.com/akimoto-junya/ouchi-hub-backend/internal/module/ouchibucket/domain/repository"

type Content interface {
}

type content struct {
	repo repository.Content
}

func NewContent(repo repository.Content) Content {
	return &content{repo: repo}
}
