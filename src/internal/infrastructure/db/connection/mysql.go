package db

import (
	commons "count_on_us/src/internal/commons/strings"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MySQL struct {
	UserName string
	Password string
	Host     string
	Port     string
	Database string
	DBType   string
}

func (m MySQL) Connect() (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		m.UserName,
		m.Password,
		m.Host,
		m.Port,
		m.Database)
	return sqlx.Open(m.DBType, connectionString)
}

func (m MySQL) IsValid() error {
	if commons.IsEmpty(m.Host) {
		return errors.New("empty host")
	}
	if commons.IsEmpty(m.UserName) {
		return errors.New("empty username")
	}
	if commons.IsEmpty(m.Port) {
		return errors.New("empty port")
	}
	if commons.IsEmpty(m.Database) {
		return errors.New("empty database name")
	}
	if commons.IsEmpty(m.Password) {
		return errors.New("empty password")
	}
	if commons.IsEmpty(m.DBType) {
		return errors.New("invalid dbtype for postgres")
	}
	return nil
}
