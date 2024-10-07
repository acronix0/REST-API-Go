package database

import (
	"database/sql"
	"fmt"

	"github.com/acronix0/REST-API-Go/internal/config"
)

type Database struct {
	db *sql.DB
}

func New(con *config.SqlConnection) (*Database, error) {
	const op = "storage.postgres.New"
	conStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		con.Host, con.Port, con.UserName, con.Password, con.Name,
	)
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Database{db}, nil
}

func (s *Database) GetDB() *sql.DB {
	return s.db
}
func (s *Database) Close() error {
	return s.db.Close()
}
