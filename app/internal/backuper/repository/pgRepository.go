package repository

import (
	"context"
	"db-backuper/config"
	"db-backuper/internal/backuper"
	"github.com/jmoiron/sqlx"
	"time"
)

type BackuperPGRepo struct {
	pgDB *sqlx.DB
	cfg  *config.Config
}

func NewBackuperPGRepo(pgDB *sqlx.DB, cfg *config.Config) backuper.PGRepo {
	return &BackuperPGRepo{pgDB: pgDB, cfg: cfg}
}

func (r *BackuperPGRepo) IsBotDBUpdated(ctx context.Context, now time.Time) (bool, error) {
	return false, nil
}
