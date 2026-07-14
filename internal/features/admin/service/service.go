package service

import (
	"N1ktarchik/Bar-Manager/internal/core/domain"
	"context"
	"log/slog"
)

type BarAdminService struct {
	repository BarAdminRepository
	log        *slog.Logger
}

type BarAdminRepository interface {
	AddCocktail(ctx context.Context, cocktail *domain.Cocktail) (*domain.Cocktail, error)
	DeleteCocktail(ctx context.Context, id int) error
}

func NewBarAdminService(repository BarAdminRepository, log *slog.Logger) *BarAdminService {
	return &BarAdminService{
		repository: repository,
		log:        log,
	}
}
