package main

import (
	"context"
	"db-backuper/config"
	usecaseBackuper "db-backuper/internal/backuper/usecase"
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

	backuperUC := usecaseBackuper.NewBackuper(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for ; true; <-ticker.C {
			now := time.Now()
			if err := backuperUC.PGBotDBBackup(ctx, now); err != nil {
				log.Println(err)
			}

			log.Printf("Successfully created backup of bot db at %v", now)
		}
	}()

	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	<-exitCh
}
