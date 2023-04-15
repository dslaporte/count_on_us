package db

import "github.com/jmoiron/sqlx"

type Database interface {
	Connect() (*sqlx.DB, error)
	IsValid() error
}
