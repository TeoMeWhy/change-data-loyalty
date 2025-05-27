package main

import (
	"flag"
	"fmt"
	"log"
	"migrator/configs"
	"migrator/controller/executer"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	duration := flag.Duration("sleep", 10, "Sleep duration between runs")
	flag.Parse()

	godotenv.Load("../../.env")
	config, err := configs.LoadConfig()
	if err != nil {
		panic(err)
	}

	dns := "%s:%s@tcp(%s:%s)/%s?parseTime=true"
	dns = fmt.Sprintf(dns, config.UserDB, config.PasswordDB, config.HostDB, config.PortDB, "points")
	dbPoints, err := sqlx.Connect("mysql", dns)
	if err != nil {
		panic(fmt.Errorf("error connecting to database: %w", err))
	}

	executer, err := executer.NewExecuter(dbPoints, config)
	if err != nil {
		panic(fmt.Errorf("error creating executer: %w", err))
	}

	log.Println("Starting migration process...")
	for {
		log.Println("Running executer...")
		if err := executer.Run(); err != nil {
			panic(fmt.Errorf("error running executer: %w", err))
		}
		log.Println("Executer run completed successfully. Waiting for the next run...")
		time.Sleep(*duration * time.Minute)
	}

}
