package collection

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"mocking-server/internal/dto/mock_server_dto/response"
	"mocking-server/internal/model"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

type Repository interface {
	InsertCollection(context.Context, *sqlx.Tx, *model.Collection) error
	SelectByWorkspaceId(context.Context, *string) (*[]response.CollectionResponse, error)
}

func (r *repository) InsertCollection(c context.Context, tx *sqlx.Tx, req *model.Collection) error {
	query := `INSERT INTO collections (id,workspace_id,reference_id,created_at, updated_at) VALUES ($1,$2,$3,$4,$5)`
	var args []interface{}

	args = append(args, req.Id, req.WorkspaceId, req.ReferenceId, req.CreatedAt, req.UpdatedAt)
	_, err := tx.ExecContext(c, query, args...)
	if err != nil {
		log.Errorf("InsertCollection.tx.ExecContext Error: %v", err)
		tx.Rollback()
	}
	return err
}
func (r *repository) SelectByWorkspaceId(c context.Context, workspaceId *string) (*[]response.CollectionResponse, error) {
	query := `SELECT cl.id,ch.name,cl.workspace_id,cl.reference_id,cl.created_at, cl.updated_at FROM collections cl JOIN childrens ch on ch.perent=cl.id  WHERE cl.workspace_id = $1`
	rows, err := r.db.QueryxContext(c, query, *workspaceId)
	if err != nil {
		return nil, err
	}
	var collections []response.CollectionResponse
	for rows.Next() {
		var collection response.CollectionResponse
		err := rows.Scan(&collection.Id, &collection.Name, &collection.WorkspaceId, &collection.ReferenceId, &collection.CreatedAt, &collection.UpdatedAt)
		if err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}
	return &collections, nil
}
