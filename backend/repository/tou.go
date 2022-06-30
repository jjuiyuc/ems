package repository

import (
	"database/sql"
)

// TOURepository godoc
type TOURepository interface {
}

type defaultTOURepository struct {
	db *sql.DB
}

// NewTOURepository godoc
func NewTOURepository(db *sql.DB) TOURepository {
	return &defaultTOURepository{db}
}
