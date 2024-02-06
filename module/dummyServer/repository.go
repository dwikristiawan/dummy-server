package dummyServer

import (
	"context"
	"database/sql"
	"dummy-server/module/dummyServer/model"
	"dummy-server/utils"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	Db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{Db: db}
}

type Repository interface {
	SetDummyRepository(context.Context, *model.DummyServer) error
	IsExist(context.Context, string) (bool, error)
	RemoveDummy(context.Context, string) error
	GetDummyByMethodAndPath(context.Context, string, string) (*[]model.DummyServer, error)
}

func (r repository) SetDummyRepository(c context.Context, dummyData *model.DummyServer) error {
	var (
		qury string
		row  *sql.Row
	)

	switch dummyData.Id {
	case "":
		qury = `INSERT INTO mock (method,path,response_conten_type, response_code, response_body, reference_id,create_at) VALUES($1,$2,$3,$4,$5,$6,$7)`
		row = r.Db.QueryRowContext(c, qury, dummyData.Method, dummyData.Path, dummyData.ResponseContenType, dummyData.ResponseCode, dummyData.ResponseBody, dummyData.CreateAt, dummyData.ReferenceId)
		break
	default:
		qury = `UPDATE mock
				SET method=$1,path=$2,response_conten_type=$3, response_code=$4, response_body=$5,update_at=$6 
				where id=$7 `
		row = r.Db.QueryRowContext(c, qury, dummyData.Method, dummyData.Path, dummyData.ResponseContenType, dummyData.ResponseCode, dummyData.ResponseBody, dummyData.UpdateAt, dummyData.Id)

	}
	return row.Err()
}
func (r repository) IsExist(c context.Context, id string) (bool, error) {
	row := r.Db.QueryRowContext(c, `select id from dummy_server where id=$1`, id)
	if row.Err() != nil {
		return false, row.Err()
	}

	var idResponse string
	err := row.Scan(&idResponse)
	if err != nil {
		return false, err
	}
	if idResponse == "" {
		return false, nil
	}
	return true, nil
}

func (r repository) RemoveDummy(c context.Context, id string) error {
	row := r.Db.QueryRowContext(c, `DELETE mock WHERE id=$1`, id)
	return row.Err()
}

func (r repository) GetDummyByMethodAndPath(c context.Context, method string, path string) (*[]model.DummyServer, error) {
	query := `SELECT response_conten_type, response_code, response_body FROM mock WHERE method=$1 and path=$2`
	row := utils.QueryRowContext(c, r.Db, query, method, path)
	var data *[]model.DummyServer
	if row.Err() != nil {
		return nil, row.Err()
	}
	err := row.Scan(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
