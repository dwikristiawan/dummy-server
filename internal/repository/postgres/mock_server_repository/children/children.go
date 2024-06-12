package children

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"mocking-server/internal/model"
)

type repository struct {
	db *sqlx.DB
}
type Repository interface {
	InsertChildren(context.Context, *sqlx.Tx, *model.Children) error
	SelectByCollectionId(context.Context, *string) (*[]model.Children, error)
	SelectByChildrenId(context.Context, *string) (*[]model.Children, error)
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}
func (r *repository) InsertChildren(c context.Context, tx *sqlx.Tx, children *model.Children) error {
	query := `INSERT INTO childrens (id,collection_id,name,perent,reference_id,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`
	var args []interface{}
	args = append(args, children.Id, children.CollectionId, children.Name, children.Perent, children.ReferenceId, children.CreatedAt, children.UpdatedAt)
	_, err := tx.Exec(query, args...)
	if err != nil {
		log.Errorf("CnsertChildren.tx.Exec Error: %v", err)
		tx.Rollback()
	}
	return err
}
func (r *repository) SelectByCollectionId(c context.Context, collectionId *string) (*[]model.Children, error) {
	query := `SELECT id,collection_id,name,perent,reference_id,created_at,updated_at FROM childrens WHERE collection_id = perent AND collection_id= $1 `
	rows, err := r.db.Queryx(query, *collectionId)
	if err != nil {
		return nil, err
	}
	var childrens []model.Children
	for rows.Next() {
		var children model.Children
		rows.Scan(&children.Id, &children.CollectionId, &children.Name, &children.Perent, &children.ReferenceId, &children.CreatedAt, &children.UpdatedAt)
		childrens = append(childrens, children)
	}
	return &childrens, nil
}
func (r *repository) SelectByChildrenId(c context.Context, childrenId *string) (*[]model.Children, error) {
	query := `SELECT id,collection_id,name,perent,reference_id,created_at,updated_at FROM childrens WHERE perent = $1`
	rows, err := r.db.Queryx(query, *childrenId)
	if err != nil {
		return nil, err
	}
	var childrens []model.Children
	for rows.Next() {
		var children model.Children
		rows.Scan(&children.Id, &children.CollectionId, &children.Name, &children.Perent, &children.ReferenceId, &children.CreatedAt, &children.UpdatedAt)
		childrens = append(childrens, children)
	}
	return &childrens, nil
}
