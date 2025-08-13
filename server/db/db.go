package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/devsrivatsa/chat_app_go-ts-react/config"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	conStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Connected to the database")
	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	if err := d.db.Close(); err != nil {
		return fmt.Errorf("error closing database: %w", err)
	}
	log.Println("Database connection closed")
	return nil
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
