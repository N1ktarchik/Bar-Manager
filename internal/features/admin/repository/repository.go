package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BarAdminRepository struct {
	pool *pgxpool.Pool
	log  *slog.Logger
}

func NewBarAdminRepository(pool *pgxpool.Pool, log *slog.Logger) *BarAdminRepository {
	return &BarAdminRepository{
		pool: pool,
		log:  log,
	}
}
