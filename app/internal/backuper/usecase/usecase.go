package usecase

import (
	"context"
	"db-backuper/config"
	"db-backuper/internal/backuper"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"time"
)

type Backuper struct {
	cfg *config.Config
}

func NewBackuper(cfg *config.Config) backuper.Backuper {
	return &Backuper{cfg: cfg}
}

func (u *Backuper) PGBotDBBackup(ctx context.Context, now time.Time) error {
	cmd := u.generateCMD()

	filePath := fmt.Sprintf(u.cfg.BackupPath,
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fileMode := os.FileMode(0600)

	outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fileMode)
	if err != nil {
		err = errors.Wrap(err, "Backuper.PGBotDBBackup.OpenFile")
		return err
	}
	defer outFile.Close()

	cmd.Stdout = outFile
	if err := cmd.Run(); err != nil {
		err = errors.Wrap(err, "Backuper.PGBotDBBackup.Run")
		return err
	}

	return nil
}

func (u *Backuper) generateCMD() *exec.Cmd {
	return exec.Command(
		"docker",
		"exec",
		"demo-db",
		"env",
		fmt.Sprintf("PGPASSWORD=%s", u.cfg.Postgres.Password),
		"pg_dump",
		"-U",
		fmt.Sprintf("%s", u.cfg.Postgres.User),
		"-d",
		fmt.Sprintf("%s", u.cfg.Postgres.DBName))
}
