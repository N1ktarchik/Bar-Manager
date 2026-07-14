package repository

import (
	"N1ktarchik/Bar-Manager/internal/core/domain"
	"N1ktarchik/Bar-Manager/internal/core/errors"
	"context"
	"log/slog"
	"time"
)

func (r *BarAdminRepository) UpdatePrice(ctx context.Context, id int, price int) (*domain.Cocktail, error) {
	r.log.Debug("new request to DB (UpdatePrice)",
		slog.Any("time:", time.Now().UTC().Format("2006-01-02 15:04:05")))

	query := `UPDATE menu SET price=$1 WHERE id=$2
			RETURNING id,name,ingridients,price`

	updatedCocktail := cocktailModel{}

	row := r.pool.QueryRow(ctx, query, price, id)

	if err := updatedCocktail.scan(row); err != nil {
		r.log.Error("error to update price of cocktail in DB", slog.Any("err", err))
		return nil, errors.INTERNAL_SERVER_ERR()
	}

	domainCocktail := modelToDomain(updatedCocktail)

	r.log.Debug("price of cocktail successfully updated in database", slog.Any("ID", updatedCocktail.Id),
		slog.Any("price", updatedCocktail.Price),
		slog.Any("time:", time.Now().UTC().Format("2006-01-02 15:04:05")))

	return &domainCocktail, nil
}
