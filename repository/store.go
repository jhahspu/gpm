package repository

import (
	"database/sql"
)

type sqliteRepository struct {
	db *sql.DB
}

func NewSqliteRepository(db *sql.DB) (SqliteRepository, error) {
	repo := &sqliteRepository{
		db: db,
	}

	return repo, nil
}
