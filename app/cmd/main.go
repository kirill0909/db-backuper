package main

import (
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

	_ = usecaseBackuper.NewBackuper(cfg)

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		for ; true; <-ticker.C {
			log.Println("hello")
		}
	}()

	exitCh := make(chan os.Signal)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	<-exitCh
}
