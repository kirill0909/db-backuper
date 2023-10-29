package usecase

import (
	"context"
	"db-backuper/config"
	"db-backuper/internal/backuper"
	"fmt"
	// "github.com/pkg/errors"
	// "os"
	"log"
	"os/exec"
	"time"
)

type BackuperUC struct {
	backuperPGRepo backuper.PGRepo
	cfg            *config.Config
}

func NewBackuperUC(backuperPGRepo backuper.PGRepo, cfg *config.Config) backuper.Usecase {
	return &BackuperUC{backuperPGRepo: backuperPGRepo, cfg: cfg}
}

func (u *BackuperUC) PGBotDBBackup(ctx context.Context, now time.Time) error {
	res, err := u.backuperPGRepo.IsBotDBUpdated(ctx, now)
	if err != nil {
		return err
	}

	if !res {
		log.Printf("Bot db was not updated at %v", now)
		return nil
	}

	// cmd := u.generateCMD()
	//
	// filePath := fmt.Sprintf(u.cfg.BackupPath,
	// 	now.Year(), now.Month(), now.Day(),
	// 	now.Hour(), now.Minute(), now.Second())
	// fileMode := os.FileMode(0600)
	//
	// outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fileMode)
	// if err != nil {
	// 	err = errors.Wrap(err, "Backuper.PGBotDBBackup.OpenFile")
	// 	return err
	// }
	// defer outFile.Close()
	//
	// cmd.Stdout = outFile
	// if err := cmd.Run(); err != nil {
	// 	err = errors.Wrap(err, "Backuper.PGBotDBBackup.Run")
	// 	return err
	// }

	return nil
}

func (u *BackuperUC) generateCMD() *exec.Cmd {
	return exec.Command(
		"docker",
		"exec",
		"boost-my-skills-boot_db_1",
		"env",
		fmt.Sprintf("PGPASSWORD=%s", u.cfg.Postgres.Password),
		"pg_dump",
		"-U",
		fmt.Sprintf("%s", u.cfg.Postgres.User),
		"-d",
		fmt.Sprintf("%s", u.cfg.Postgres.DBName))
}
