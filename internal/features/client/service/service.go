package service

import (
	"N1ktarchik/Bar-Manager/internal/core/domain"
	"context"
	"log/slog"
)

type BarClientService struct {
	repository BarClientRepository
}

type BarClientRepository interface {
	GetCocktails(ctx context.Context) ([]domain.Cocktail, error)
}

func NewBarClientService(repository BarClientRepository, log *slog.Logger) *BarClientService {
	return &BarClientService{
		repository: repository,
	}
}
