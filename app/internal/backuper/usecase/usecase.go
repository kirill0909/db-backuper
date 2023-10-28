package usecase

import (
	"context"
	"db-backuper/config"
	"db-backuper/internal/backuper"
	"log"
)

type Backuper struct {
	cfg *config.Config
}

func NewBackuper(cfg *config.Config) backuper.Backuper {
	return &Backuper{cfg: cfg}
}

func (u *Backuper) PGBackup(ctx context.Context) {
	log.Println("Hello from backuper usecase")
}
