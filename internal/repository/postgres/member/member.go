package member

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
	InsertMember(context.Context, *sqlx.Tx, *model.Member) error
}

func NewRepository(db *sqlx.DB) Reppsitory {
	return &repository{
		db: db,
	}
}
func (repo repository) InsertMember(c context.Context, tx *sqlx.Tx, req *model.Member) error {
	query := `
	INSERT INTO member(
		id,
		type_id,
		user_id,
		access,
		is_acctive,
		created_at,
		updated_at
	)
	values(
		$1,$2,$3,$4,$5,$6,$7,$8
	)`
	var args []interface{}
	args = append(args,
		req.Id,
		req.UserId,
		req.Access,
		req.IsActive,
		req.CreatedAt,
		req.UpdatedAt)
	_, err := tx.ExecContext(c, query, args...)
	if err != nil {
		tx.Rollback()
		log.Errorf("InsertMember.repo.db.DB.ExecContext Err: %v", err)
		return err
	}
	return nil
}
