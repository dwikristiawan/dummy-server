package mockdata

import (
	"context"
	"mocking-server/internal/model"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type repository struct {
	db *sqlx.DB
}
type Reppsitory interface {
	DBBegin() (*sqlx.Tx, error)
	InsertMockData(context.Context, *sqlx.Tx, *model.MockData) error
	SelectMockData(context.Context, *model.MockData) (*[]model.MockData, error)
}

func NewRepository(db *sqlx.DB) Reppsitory {
	return &repository{db: db}
}
func (repo repository) DBBegin() (*sqlx.Tx, error) {
	tx, err := repo.db.Beginx()
	return tx, err
}

func (repo repository) InsertMockData(c context.Context, tx *sqlx.Tx, req *model.MockData) error {
	var args []interface{}
	query := `INSERT INTO mock_data(
		id,
		children_id,
		request_method,
		path,
		request_header,
		response_header,
		request_body,
		response_body,
		response_code,
		reference_id,
		created_at,
		updated_at
	)values(
		$1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12
	)`

	req.RequestHeader = nil
	args = append(args,
		req.Id,
		req.ChildrenId,
		req.RequestMethod,
		req.Path, req.RequestHeader)
	// if req.RequestHeader == nil {
	// 	args = append(args, "{}")
	// } else {
	// 	args = append(args, req.RequestHeader)
	// }
	args = append(args,
		req.ResponseHeader,
		req.RequestBody,
		req.ResponseBody,
		req.ResponseCode,
		req.ReferenceId,
		req.CreatedAt,
		req.UpdatedAt)
	_, err := tx.ExecContext(c, query, args...)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (repo repository) SelectMockData(c context.Context, req *model.MockData) (*[]model.MockData, error) {
	builder := squirrel.Select(
		`id,
	children_id,
	request_method,
	path,
	request_header,
	response_header,
	request_body,
	response_body,
	response_code,
	reference_id,
	created_at,
	updated_at`).From(`mock_data`).PlaceholderFormat(squirrel.Dollar)
	if req.Id != "" {
		builder = builder.Where(squirrel.Eq{"id": req.Id})
	}
	if req.RequestMethod != "" {
		builder = builder.Where(squirrel.Eq{"request_method": req.RequestMethod})
	}
	if req.Path != "" {
		builder = builder.Where(squirrel.Eq{"path": req.Path})
	}
	if req.RequestBody == "" {
		builder = builder.Where(squirrel.Eq{"request_body": req.RequestBody})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		log.Errorf("SelectMockData.builder.ToSql() Err: %v", err)
		return nil, err
	}
	rows, err := repo.db.QueryContext(c, query, args...)
	if err != nil {
		log.Errorf("SelectMockData.repo.db.QueryContext Err: %v", err)
		return nil, err
	}
	var mockDatas []model.MockData
	for rows.Next() {
		var mockData model.MockData
		err = rows.Scan(
			&mockData.Id,
			&mockData.ChildrenId,
			&mockData.RequestMethod,
			&mockData.Path,
			&mockData.RequestHeader,
			&mockData.ResponseHeader,
			&mockData.RequestBody,
			&mockData.ResponseBody,
			&mockData.ResponseCode,
			&mockData.ReferenceId,
			&mockData.CreatedAt,
			&mockData.UpdatedAt,
		)
		mockDatas = append(mockDatas, mockData)
	}
	return &mockDatas, nil
}
