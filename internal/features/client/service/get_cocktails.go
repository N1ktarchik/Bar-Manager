package service

import (
	"N1ktarchik/Bar-Manager/internal/core/domain"
	"context"
)

func (s *BarClientService) GetCocktails(ctx context.Context) ([]domain.Cocktail, error) {
	return s.repository.GetCocktails(ctx)
}
