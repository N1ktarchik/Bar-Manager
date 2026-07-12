package transport

import (
	"context"
	"hello/Desktop/pet-projects/bar-manager/internal/core/domain"
	"log/slog"
)

type BarClientHandlerHTTP struct {
	service BarClientService
	log     *slog.Logger
}

type BarClientService interface {
	GetCocktails(ctx context.Context) ([]domain.Cocktail, error)
}

func NewBarTransportHTTP(service BarClientService, log *slog.Logger) *BarClientHandlerHTTP {
	return &BarClientHandlerHTTP{
		service: service,
		log:     log,
	}
}
