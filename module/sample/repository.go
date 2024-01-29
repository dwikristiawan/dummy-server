package sample

import "github.com/jmoiron/sqlx"

type repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{Db: db}
}

type Repository interface {
}
