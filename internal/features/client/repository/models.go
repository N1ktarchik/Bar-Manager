package repository

import (
	"N1ktarchik/Bar-Manager/internal/core/domain"

	"github.com/jackc/pgx/v5"
)

type cocktailModel struct {
	Id          int
	Name        string
	Ingridients string
	Price       int
}

func (m *cocktailModel) scan(row pgx.Row) error {
	return row.Scan(
		&m.Id,
		&m.Name,
		&m.Ingridients,
		&m.Price,
	)
}

func modelToDomain(model cocktailModel) domain.Cocktail {
	return domain.Cocktail{
		Id:          model.Id,
		Name:        model.Name,
		Ingridients: model.Ingridients,
		Price:       model.Price,
	}
}
