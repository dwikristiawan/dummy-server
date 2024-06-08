package workspace

import (
	"context"
	"mocking-server/internal/model"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type repository struct {
	db *sqlx.DB
}

type Reppsitory interface {
	InsertWorkSpace(context.Context, *sqlx.Tx, *model.WorkSpace) error
	DBBegin() (*sqlx.Tx, error)
}

func NewRepository(db *sqlx.DB) Reppsitory {
	return &repository{
		db: db,
	}
}

func (repo repository) DBBegin() (*sqlx.Tx, error) {
	tx, err := repo.db.Beginx()
	return tx, err
}
func (repo repository) InsertWorkSpace(c context.Context, tx *sqlx.Tx, req *model.WorkSpace) error {
	query :=
		`INSERT INTO work_space(
		 id,
		 name,
		 reference_id,
		 created_at,
		 updated_at
		)
		values(
			$1,
			$2,
			$3,		
			$4,
			$5
		)`
	var args []interface{}
	args = append(args,
		req.Id,
		req.Name,
		req.ReferenceId,
		req.CreatedAt,
		req.UpdatedAt)

	_, err := tx.ExecContext(c, query, args...)
	if err != nil {
		tx.Rollback()
		log.Errorf("InsertWorkSpace.repo.db.DB.ExecContext Err: %v", err)
		return err
	}
	return nil
}
