package usecase

import (
	"context"
	"db-backuper/config"
	"db-backuper/internal/backuper"
	"log"
	"time"
	// "os/exec"
	"fmt"
	"os"
)

type Backuper struct {
	cfg *config.Config
}

func NewBackuper(cfg *config.Config) backuper.Backuper {
	return &Backuper{cfg: cfg}
}

func (u *Backuper) PGBotDBBackup(ctx context.Context) error {
	// cmd := exec.Command("docker", "exec", "demo-db", "env", "PGPASSWORD=123", "pg_dump", "-U", "postgres", "-d", "demo")

	t := time.Now()
	fileName := fmt.Sprintf("bot_db_backup_%d_%d_%d_%d.sql", t.Day(), t.Hour(), t.Minute(), t.Second())
	outFile, err := os.Create(fileName)

	return nil
}
