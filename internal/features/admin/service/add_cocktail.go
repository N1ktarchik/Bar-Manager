package service

import (
	"N1ktarchik/Bar-Manager/internal/core/domain"
	"N1ktarchik/Bar-Manager/internal/core/errors"
	"log/slog"

	"context"
)

func (s *BarAdminService) AddCocktail(ctx context.Context, cocktail *domain.Cocktail) (*domain.Cocktail, error) {
	if len(cocktail.Name) < 3 {
		s.log.Debug(*errors.SHORT_NAME_ERR().Msg(), slog.String("Name", cocktail.Name))
		return nil, errors.SHORT_NAME_ERR()
	}

	if cocktail.Price <= 0 {
		s.log.Debug(*errors.INVALID_PRICE_ERR().Msg(), slog.Int("Price", cocktail.Price))
		return nil, errors.INVALID_PRICE_ERR()
	}

	savedCocktail, err := s.repository.AddCocktail(ctx, cocktail)
	if err != nil {
		return nil, err
	}

	return savedCocktail, nil
}
