package transport

import (
	"N1ktarchik/Bar-Manager/internal/core/domain"
	"context"
	"log/slog"
)

type BarAdminHandlerHTTP struct {
	barService  BarAdminService
	authService AuthService
	log         *slog.Logger
}

type BarAdminService interface {
	AddCocktail(ctx context.Context, cocktail *domain.Cocktail) (*domain.Cocktail, error)
	UpdatePrice(ctx context.Context, id string, price int) (*domain.Cocktail, error)
	DeleteCocktail(ctx context.Context, id string) error
}

type AuthService interface {
	CreateJWT(password string) (string, error)
}

func NewBarAdminTransportHTTP(barService BarAdminService, authService AuthService, log *slog.Logger) *BarAdminHandlerHTTP {
	return &BarAdminHandlerHTTP{
		barService:  barService,
		authService: authService,
		log:         log,
	}
}
