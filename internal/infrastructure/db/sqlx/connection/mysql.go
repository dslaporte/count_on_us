package db

import (
	pkg_strings "count_on_us/pkg/strings"
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
	if pkg_strings.IsEmpty(m.Host) {
		return errors.New("empty host")
	}
	if pkg_strings.IsEmpty(m.UserName) {
		return errors.New("empty username")
	}
	if pkg_strings.IsEmpty(m.Port) {
		return errors.New("empty port")
	}
	if pkg_strings.IsEmpty(m.Database) {
		return errors.New("empty database name")
	}
	if pkg_strings.IsEmpty(m.Password) {
		return errors.New("empty password")
	}
	if pkg_strings.IsEmpty(m.DBType) {
		return errors.New("invalid dbtype for postgres")
	}
	return nil
}
