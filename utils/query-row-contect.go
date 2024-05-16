package utils

// import (
// 	"context"

// 	"github.com/jmoiron/sqlx"
// 	"github.com/labstack/gommon/log"
// )

// func QueryRowContext(c context.Context, db *sqlx.DB, model interface{}, query string, [

// ]interface{}) (interface{}, error) {
// 	row := db.QueryRowContext(c, query, args...)
// 	if row.Err() != nil {
// 		log.Errorf("dummy-server repository queryDb, error: ", row.Err())
// 		return nil, row.Err()
// 	}
// 	if model != nil {
// 		err := row.Scan(&model)
// 		if err != nil {
// 			log.Errorf("dummy-server repository row.Scan, error: ", err)
// 			return nil, err
// 		}
// 		return model, nil
// 	}
// 	return row, nil

// }
