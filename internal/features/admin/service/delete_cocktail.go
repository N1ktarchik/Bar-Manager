package service

import (
	"N1ktarchik/Bar-Manager/internal/core/errors"
	"context"
	"log/slog"
	"strconv"
)

func (s *BarAdminService) DeleteCocktail(ctx context.Context, id string) error {
	parsedID, err := strconv.Atoi(id)

	if err != nil {
		s.log.Debug("Parsing id failed", slog.String("string id", id))
		return errors.BAD_REQUEST_ERR()
	}

	if parsedID <= 0 {
		s.log.Debug(*errors.INVALID_ID_ERR().Msg(), slog.Int("ID", parsedID))
		return errors.INVALID_ID_ERR()
	}

	return s.repository.DeleteCocktail(ctx, parsedID)
}
