package types

import (
	"context"
	"mocking-server/internal/model"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}
type Reppsitory interface {
	InsertTypes(context.Context, *sqlx.Tx, model.Types) error
}

func NewTypes(db *sqlx.DB) Reppsitory {
	return &repository{db: db}
}

func (repo repository) InsertTypes(c context.Context, tx *sqlx.Tx, req model.Types) error {
	return nil
}
