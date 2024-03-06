package usecase

import (
	"context"
	"db-backuper/app/config"
	"db-backuper/app/internal/backuper"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/pkg/errors"
)

type BackuperUC struct {
	backuperPGRepo backuper.PGRepo
	cfg            *config.Config
	checkTime      int64
}

func NewBackuperUC(backuperPGRepo backuper.PGRepo, cfg *config.Config) backuper.Usecase {
	return &BackuperUC{backuperPGRepo: backuperPGRepo, cfg: cfg, checkTime: 0}
}

func (u *BackuperUC) PGBotDBBackup(ctx context.Context, now time.Time) error {

	res, err := u.backuperPGRepo.IsBotDBUpdated(ctx, u.checkTime)
	if err != nil {
		return err
	}
	u.checkTime = now.Unix()

	switch res {
	case true:
		log.Printf("Bot db was updated at %v", now)
		if err := u.handleUpdatedCase(now); err != nil {
			return err
		}
	case false:
		log.Printf("Bot db was not updated at %v", now)
		return nil
	}

	return nil
}

func (u *BackuperUC) handleUpdatedCase(now time.Time) error {
	cmd := u.generateCMD()

	filePath := fmt.Sprintf(u.cfg.BotDBBackupPath,
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

func (u *BackuperUC) generateCMD() *exec.Cmd {
	return exec.Command(
		"docker",
		"exec",
		u.cfg.BotDBContainerName,
		"env",
		fmt.Sprintf("PGPASSWORD=%s", u.cfg.Postgres.Password),
		"pg_dump",
		"-U",
		fmt.Sprintf("%s", u.cfg.Postgres.User),
		"-d",
		fmt.Sprintf("%s", u.cfg.Postgres.DBName))
}
