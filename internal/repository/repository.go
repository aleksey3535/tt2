package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

var ErrNotFound = errors.New("subscription not found")

type Repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}