package repository

import (
	"N1ktarchik/Bar-Manager/internal/core/errors"
	"context"
	"log/slog"
)

func (r *BarAdminRepository) DeleteCocktail(ctx context.Context, id int) error {
	r.log.Debug("new request to DB (DeleteCocktail)")

	query := `DELETE FROM menu WHERE id=$1`

	result, err := r.pool.Exec(ctx, query, id)

	if err != nil {
		r.log.Error("error to delete cocktail in DB", slog.Any("err", err))
		return errors.INTERNAL_SERVER_ERR()
	}

	if result.RowsAffected() == 0 {
		r.log.Warn("cocktail not faund", slog.Any("ID", id))
		return errors.ID_NOT_FAUND_ERR()
	}

	r.log.Debug("cocktail successfully deleted in database")

	return nil
}
