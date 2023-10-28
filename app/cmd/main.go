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
			if err := backuperUC.PGBackup(ctx); err != nil {
				log.Println(err)
			}
		}
	}()

	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	<-exitCh
}
