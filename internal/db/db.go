package db

import (
	"EasyFinGo/internal/config"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)


var DB *sql.DB


func Init() {
	cfg := config.Load()

	var err error 
	DB, err = sql.Open("pgx", cfg.PostgresURL)

	if err != nil {
		log.Fatalf("Unable to open database connection: %v\n", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Cannot ping database: %v\n", err)
	}

	log.Println("Database connection established")
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}