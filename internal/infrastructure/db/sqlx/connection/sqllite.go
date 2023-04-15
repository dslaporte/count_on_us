package db

import (
	pkg_strings "count_on_us/pkg/strings"
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
	if pkg_strings.IsEmpty(s.Database) {
		return errors.New("empty database")
	}
	if strings.TrimSpace(s.Database) != ":memory:" {
		if !fs.ValidPath(s.Database) {
			return errors.New("invalid path for sqlite database")
		}
	}
	if pkg_strings.IsEmpty(s.DBtype) {
		return errors.New("invalid db type for sqlite3")
	}
	return nil
}
