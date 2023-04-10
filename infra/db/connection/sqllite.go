package db

import (
	commons "count_on_us/internal/commons/strings"
	"errors"
	"io/fs"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	Database string
	DBtype   string
}

func (s SQLite) Connect() (*sqlx.DB, error) {
	return sqlx.Open(s.DBtype, s.Database)
}

func (s SQLite) IsValid() error {
	if commons.IsEmpty(s.Database) {
		return errors.New("empty database")
	}
	if strings.TrimSpace(s.Database) != ":memory:" {
		if !fs.ValidPath(s.Database) {
			return errors.New("invalid path for sqlite database")
		}
	}
	if commons.IsEmpty(s.DBtype) {
		return errors.New("invalid db type for sqlite3")
	}
	return nil
}
