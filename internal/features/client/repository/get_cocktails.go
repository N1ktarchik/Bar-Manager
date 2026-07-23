package repository

import (
	"N1ktarchik/Bar-Manager/internal/core/domain"
	"N1ktarchik/Bar-Manager/internal/core/errors"
	"context"
	"log/slog"
)

func (r *BarClientRepository) GetCocktails(ctx context.Context) ([]domain.Cocktail, error) {
	r.log.Debug("new request to DB (GetCocktails)")

	query := `SELECT id,name,ingridients,price FROM menu`

	rows, err := r.pool.Query(ctx, query)
	cocktails := make([]domain.Cocktail, 0, 10)

	for rows.Next() {
		var c cocktailModel

		if err = c.scan(rows); err != nil {
			r.log.Error("error to scan cocktails from DB", slog.Any("err", err))
			return nil, errors.INTERNAL_SERVER_ERR()
		}

		domainCocktail := modelToDomain(c)
		cocktails = append(cocktails, domainCocktail)
	}

	if err := rows.Err(); err != nil {
		r.log.Error("rows cursor error", slog.Any("err", err))
		return nil, errors.INTERNAL_SERVER_ERR()
	}

	r.log.Debug("All cocktails from the database have been successfully retrieved")
	return cocktails, nil
}
