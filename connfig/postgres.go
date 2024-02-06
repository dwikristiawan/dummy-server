package config

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     int    `envconfig:"POSTGRES_PORT" required:"true"`
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Dbname   string `envconfig:"POSTGRES_DATABASE" required:"true"`

	MaxConnectionLifetime time.Duration `envconfig:"DB_MAX_CONN_LIFE_TIME" default:"300s"`
	MaxOpenConnection     int           `envconfig:"DB_MAX_OPEN_CONNECTION" default:"100"`
	MaxIdleConnection     int           `envconfig:"DB_MAX_IDLE_CONNECTION" default:"10"`
}

func (p Postgres) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Dbname)
}

func OpenPostgresDatabaseConnection(pg Postgres) (*sqlx.DB, error) {

	dbConn, err := sqlx.Connect("postgres", pg.ConnectionString())
	if err != nil {
		log.Errorf("Err > %v", err)
	}

	dbConn.SetConnMaxLifetime(pg.MaxConnectionLifetime)
	dbConn.SetMaxOpenConns(pg.MaxOpenConnection)
	dbConn.SetMaxIdleConns(pg.MaxIdleConnection)

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}
