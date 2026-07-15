package repository

import (
	"N1ktarchik/Bar-Manager/internal/core/domain"
	"N1ktarchik/Bar-Manager/internal/core/errors"

	"context"
	"log/slog"
)

func (r *BarAdminRepository) AddCocktail(ctx context.Context, cocktail *domain.Cocktail) (*domain.Cocktail, error) {
	r.log.Debug("new request to DB (AddCocktail)")

	query := `INSERT INTO menu (name,ingridients,price) 
			VALUES ($1,$2,$3) 
			RETURNIG id,name,ingridients,price`

	savedCocktail := cocktailModel{}

	row := r.pool.QueryRow(ctx, query, cocktail.Name, cocktail.Ingridients, cocktail.Price)

	if err := savedCocktail.scan(row); err != nil {
		r.log.Error("error to crerate new cocktail in DB", slog.Any("err", err))
		return nil, errors.INTERNAL_SERVER_ERR()
	}

	domainCocktail := modelToDomain(savedCocktail)

	r.log.Debug("cocktail successfully created in database", slog.Any("ID", savedCocktail.Id))
	return &domainCocktail, nil
}
