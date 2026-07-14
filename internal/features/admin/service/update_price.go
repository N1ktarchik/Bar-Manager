package service

import (
	"N1ktarchik/Bar-Manager/internal/core/domain"
	"N1ktarchik/Bar-Manager/internal/core/errors"
	"context"
	"log/slog"
	"strconv"
)

func (s *BarAdminService) UpdatePrice(ctx context.Context, id string, price int) (*domain.Cocktail, error) {
	parsedID, err := strconv.Atoi(id)

	if err != nil {
		s.log.Debug("Parsing id failed", slog.String("string id", id))
		return nil, errors.BAD_REQUEST_ERR()
	}

	if parsedID <= 0 {
		s.log.Debug(*errors.INVALID_ID_ERR().Msg(), slog.Int("ID", parsedID))
		return nil, errors.INVALID_ID_ERR()
	}

	if price <= 0 {
		s.log.Debug(*errors.INVALID_PRICE_ERR().Msg(), slog.Int("Price", price))
		return nil, errors.INVALID_PRICE_ERR()
	}

	updatedCocktail, err := s.repository.UpdatePrice(ctx, parsedID, price)
	if err != nil {
		return nil, err
	}

	return updatedCocktail, nil
}
