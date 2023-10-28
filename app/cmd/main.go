package main

import (
	"context"
	"db-backuper/app/config"
	"db-backuper/app/pkg/storage/postgres"
	"log"
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

	ctx := context.Background()
	pgDB, err := postgres.InitPGDB(ctx, cfg)
	if err != nil {
		log.Fatalf("PostgreSQL init error: %s", err.Error())
	} else {
		log.Printf("PostgreSQL status connection: %#v", pgDB.Stats())
	}

	log.Println(pgDB)

}
