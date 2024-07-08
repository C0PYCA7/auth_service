package main

import (
	"auth_service/internal/config"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {

	cfg := config.LoadConfig()
	path := "././internal/migration"
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbConfig.Host, cfg.DbConfig.Port, cfg.DbConfig.Username, cfg.DbConfig.Password, cfg.DbConfig.Database,
	)
	db, err := goose.OpenDBWithDriver("pgx", connectionString)
	if err != nil {
		log.Printf("Failed to open database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	err = goose.Up(db, path)
	if err != nil {
		log.Printf("Failed to apply migration: %v\n", err)
		os.Exit(1)
	}
}
