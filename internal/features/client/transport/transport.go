package transport

import (
	"N1ktarchik/Bar-Manager/internal/core/domain"
	"context"
	"log/slog"
)

type BarClientHandlerHTTP struct {
	service BarClientService
	log     *slog.Logger
}

type BarClientService interface {
	GetCocktails(ctx context.Context) ([]domain.Cocktail, error)
}

func NewBarClientTransportHTTP(service BarClientService, log *slog.Logger) *BarClientHandlerHTTP {
	return &BarClientHandlerHTTP{
		service: service,
		log:     log,
	}
}
