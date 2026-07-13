package transport

import (
	"N1ktarchik/Bar-Manager/internal/core/domain"
	"context"
	"log/slog"
)

type BarAdminHandlerHTTP struct {
	service BarAdminService
	log     *slog.Logger
}

type BarAdminService interface {
	AddCocktail(ctx context.Context, cocktail *domain.Cocktail) (*domain.Cocktail, error)
	UpdatePrice(ctx context.Context, id string, price int) (*domain.Cocktail, error)
	DeleteCocktail(ctx context.Context, id string) error
}

func NewBarTransportHTTP(service BarAdminService, log *slog.Logger) *BarAdminHandlerHTTP {
	return &BarAdminHandlerHTTP{
		service: service,
		log:     log,
	}
}
