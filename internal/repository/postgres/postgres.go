package postgres

import "github.com/jmoiron/sqlx"

type repository struct {
	db *sqlx.DB
}
type Repository interface {
	DBBegin() (*sqlx.Tx, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (repo repository) DBBegin() (*sqlx.Tx, error) {
	tx, err := repo.db.Beginx()
	return tx, err
}
