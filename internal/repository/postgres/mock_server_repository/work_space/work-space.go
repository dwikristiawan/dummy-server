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
	SelectWorkSpaceByMemberId(context.Context, *string) (*[]model.WorkSpace, error)
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
		`INSERT INTO workspaces(
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
func (repo repository) SelectWorkSpaceByMemberId(c context.Context, id *string) (*[]model.WorkSpace, error) {
	query := `SELECT ws.id,ws.name,ws.reference_id,ws.created_at,ws.updated_at FROM workspaces ws join members m on m.workspace_id = ws.id WHERE  m.user_id= $1`
	rows, err := repo.db.Queryx(query, *id)
	if err != nil {
		return nil, err
	}
	var workspaces []model.WorkSpace
	for rows.Next() {
		var workSpace model.WorkSpace
		rows.StructScan(&workSpace)
		workspaces = append(workspaces, workSpace)
	}
	return &workspaces, nil
}
