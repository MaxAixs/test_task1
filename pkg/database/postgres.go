package database

import (
	"database/sql"
	"fmt"
	"log"
)

type DBCfg struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg DBCfg) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Printf("error connect to DB: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("error ping DB: %v", err)
		return nil, err
	}

	return db, nil
}
