package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/anti-duhring/goncurrency/pkg/logger"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Init() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	logger.Debug(connStr)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		logger.Error("failed to connect to database", err)
		return nil, err
	}

	return db, nil
}
