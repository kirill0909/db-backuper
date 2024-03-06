package main

import (
	"context"
	"db-backuper/app/config"
	repositoryBackuper "db-backuper/app/internal/backuper/repository"
	usecaseBackuper "db-backuper/app/internal/backuper/usecase"
	"db-backuper/app/pkg/storage/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfgFile, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf(err.Error())
	}
	log.Println("Config loaded")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pgDB, err := postgres.InitPGDB(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("PostgreSQL connection stats: %#v", pgDB.Stats())
	}

	backuperPGRepo := repositoryBackuper.NewBackuperPGRepo(pgDB, cfg)
	backuperUC := usecaseBackuper.NewBackuperUC(backuperPGRepo, cfg)

	go func() {
		ticker := time.NewTicker(time.Second * 15)
		for ; true; <-ticker.C {
			now := time.Now()
			if err := backuperUC.PGBotDBBackup(ctx, now); err != nil {
				log.Println(err)
			}
		}
	}()

	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	<-exitCh
}
