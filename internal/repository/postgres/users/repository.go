package users

import (
	"context"
	"errors"
	"mocking-server/internal/model"
	"mocking-server/utils"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

type Repository interface {
	SelectUser(context.Context, *model.Users) (*[]model.Users, error)
	InsertUser(context.Context, *model.Users) error
	UpdateUser(context.Context, *model.Users) error
	DeleteUser(context.Context, *model.Users) error
}

func (repo repository) SelectUser(c context.Context, req *model.Users) (*[]model.Users, error) {
	builder := squirrel.Select(`*`).From(`users`).PlaceholderFormat(squirrel.Dollar)
	if req.Id != "" {
		builder = builder.Where(squirrel.Eq{`id`: req.Id})
	}
	if req.Username != "" {
		builder = builder.Where(squirrel.Eq{`username`: req.Username})
	}
	if req.Name != "" {
		builder = builder.Where(squirrel.Eq{`name`: req.Name})
	}
	if req.Status != "" {
		builder = builder.Where(squirrel.Eq{`status`: req.Status})
	}
	query, args, err := builder.ToSql()
	if err != nil {
		log.Errorf("Err SelectUser.builder.ToSql() Err > %v", err)
		return nil, err
	}

	rows, err := repo.db.QueryContext(c, query, args...)
	if err != nil {
		log.Errorf("Err SelectUser.QueryxContext Err > %v", err)
		return nil, err
	}

	users := make([]model.Users, 0)
	for rows.Next() {
		var user model.Users
		rows.Scan(
			&user.Id,
			&user.Username,
			&user.Name,
			&user.Password,
			&user.Status,
			&user.Roles,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		users = append(users, user)

	}
	return &users, nil
}
func (repo repository) InsertUser(c context.Context, req *model.Users) error {
	var argTmp []interface{}

	curent := time.Now()
	req.CreatedAt = &curent
	req.Id = utils.IdUuid()
	query := `INSERT INTO users (
		id,
		username,
		name,
		password,
		roles,
		status,
		created_at,
		updated_at
		)
		values(
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8
		)`
	argTmp = append(argTmp, req.Id)
	argTmp = append(argTmp, req.Username)
	argTmp = append(argTmp, req.Name)
	argTmp = append(argTmp, req.Password)
	argTmp = append(argTmp, req.Roles)
	argTmp = append(argTmp, req.Status)
	argTmp = append(argTmp, req.CreatedAt)
	argTmp = append(argTmp, req.UpdatedAt)

	_, err := repo.db.ExecContext(c, query, argTmp...)
	if err != nil {
		log.Errorf("Err InsertCollection.repo.db.ExecContext Err > %v", err)
		return err
	}
	return nil
}
func (repo repository) UpdateUser(c context.Context, req *model.Users) error {

	builder := squirrel.Update(`users`).Where(squirrel.Eq{`id`: req.Id}).PlaceholderFormat(squirrel.Dollar)

	if req.Username != "" {
		builder = builder.Set(`username`, req.Username)
	}
	if req.Name != "" {
		builder = builder.Set(`name`, req.Name)
	}
	if req.Password != "" {
		builder = builder.Set(`password`, req.Password)
	}
	if req.Status != "" {
		builder = builder.Set(`status`, req.Status)
	}
	builder = builder.Set(`updated_at`, utils.IdUuid())

	query, args, err := builder.ToSql()
	if err != nil {
		log.Errorf("Err InsertCollection.builder.ToSql Err > %v", err)
		return err
	}
	result, err := repo.db.ExecContext(c, query, args...)
	if err != nil {
		return err
	}
	if rowsEfect, _ := result.RowsAffected(); rowsEfect > 0 {
		err := errors.New("noting insert")
		log.Errorf("Err InserUser.ExecContext Err > %v", err)
		return err
	}
	return nil
}
func (repo repository) DeleteUser(c context.Context, req *model.Users) error {
	builder := squirrel.Delete(`users`).PlaceholderFormat(squirrel.Dollar)
	if req.Id != "" {
		builder = builder.Where(squirrel.Eq{`id`: req.Id})
	}
	if req.Username != "" {
		builder = builder.Where(squirrel.Eq{`username`: req.Username})
	}
	if req.Name != "" {
		builder = builder.Where(squirrel.Eq{`name`: req.Name})
	}
	if req.Status != "" {
		builder = builder.Where(squirrel.Eq{`status`: req.Status})
	}
	query, args, err := builder.ToSql()
	if err != nil {
		log.Errorf("Err DeleteUser.builder.ToSql Err > %v", err)
		return err
	}
	result, err := repo.db.ExecContext(c, query, args...)
	strErr := ""
	if err != nil {
		strErr = strErr + err.Error()
	}
	if rowEfect, _ := result.RowsAffected(); rowEfect == 0 {
		strErr = strErr + "row efect 0"
	}
	if strErr != "" {
		err = errors.New(strErr)
		log.Errorf("Err DeleteUser.ExecContext Err > %v", err)
		return err
	}
	return nil
}
