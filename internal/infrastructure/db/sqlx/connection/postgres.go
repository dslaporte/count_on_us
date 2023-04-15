package db

import (
	pkg_strings "count_on_us/pkg/strings"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgreSQL struct {
	Host     string
	Port     string
	UserName string
	Password string
	Database string
	SSLMode  string
	DBType   string
}

func (p PostgreSQL) Connect() (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.UserName, p.Password, p.Database, p.SSLMode)
	return sqlx.Open(string(p.DBType), connectionString)
}

func (p PostgreSQL) IsValid() error {
	if pkg_strings.IsEmpty(p.Host) {
		return errors.New("empty host")
	}
	if pkg_strings.IsEmpty(p.UserName) {
		return errors.New("empty username")
	}
	if pkg_strings.IsEmpty(p.Port) {
		return errors.New("empty port")
	}
	if pkg_strings.IsEmpty(p.Database) {
		return errors.New("empty database name")
	}
	if pkg_strings.IsEmpty(p.Password) {
		return errors.New("empty password")
	}
	if pkg_strings.IsEmpty(p.DBType) {
		return errors.New("invalid dbtype for postgres")
	}
	return nil
}
