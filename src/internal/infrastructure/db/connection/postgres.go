package db

import (
	commons "count_on_us/src/internal/commons/strings"
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
	if commons.IsEmpty(p.Host) {
		return errors.New("empty host")
	}
	if commons.IsEmpty(p.UserName) {
		return errors.New("empty username")
	}
	if commons.IsEmpty(p.Port) {
		return errors.New("empty port")
	}
	if commons.IsEmpty(p.Database) {
		return errors.New("empty database name")
	}
	if commons.IsEmpty(p.Password) {
		return errors.New("empty password")
	}
	if commons.IsEmpty(p.DBType) {
		return errors.New("invalid dbtype for postgres")
	}
	return nil
}
