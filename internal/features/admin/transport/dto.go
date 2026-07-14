package transport

import "N1ktarchik/Bar-Manager/internal/core/domain"

type cocktailDTO struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Ingridients string `json:"ingridients"`
	Price       int    `json:"price"`
}

type authDTO struct {
	Password string `json:"password"`
}

func (c *cocktailDTO) ToDomain() *domain.Cocktail {
	return &domain.Cocktail{
		Id:          c.Id,
		Name:        c.Name,
		Ingridients: c.Ingridients,
		Price:       c.Price,
	}
}
