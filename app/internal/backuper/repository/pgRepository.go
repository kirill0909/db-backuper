package repository

import (
	"context"
	"db-backuper/app/config"
	"db-backuper/app/internal/backuper"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type BackuperPGRepo struct {
	db  *sqlx.DB
	cfg *config.Config
}

func NewBackuperPGRepo(db *sqlx.DB, cfg *config.Config) backuper.PGRepo {
	return &BackuperPGRepo{db: db, cfg: cfg}
}

func (r *BackuperPGRepo) IsBotDBUpdated(ctx context.Context, now int64) (bool, error) {

	var result bool
	if err := r.db.GetContext(ctx, &result, queryIsBotDBUpdated, now); err != nil {
		err = errors.Wrap(err, "BackuperPGRepo.IsBotDBUpdated.queryIsBotDBUpdated")
		return false, err
	}

	return result, nil
}
