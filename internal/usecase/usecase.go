package usecase

import (
	"github.com/rizkyrsyd28/recloth-backend/internal/repository"
)

type Usecase interface {
	AuthUsecase
	ItemUsecase
}

type usecase struct {
	repo repository.Repo
}

func NewUsecase(r repository.Repo) usecase {
	return usecase{repo: r}
}
