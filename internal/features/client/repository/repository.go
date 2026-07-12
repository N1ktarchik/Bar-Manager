package repository

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BarClientRepository struct {
	pool *pgxpool.Pool
	log  *slog.Logger
}

func NewTasksRepository(pool *pgxpool.Pool, log *slog.Logger) *TasksRepository {
	return &TasksRepository{
		pool: pool,
		log:  log,
	}
}
