package db

import (
	"count_on_us/configs"
	"fmt"

	"github.com/jmoiron/sqlx"
)

var database_types = []string{
	"postgres",
	"mysql",
	"sqlite3",
}

func NewDatabaseFactory(cfg *configs.DatabaseConfig) (*sqlx.DB, error) {
	switch cfg.DBType {
	case "postgres":
		{
			postgresConnection := PostgreSQL{
				Host:     cfg.Host,
				Port:     cfg.Port,
				UserName: cfg.Username,
				Password: cfg.Password,
				Database: cfg.Database,
				SSLMode:  cfg.SSLMode,
				DBType:   string(cfg.DBType),
			}
			if err := postgresConnection.IsValid(); err != nil {
				return nil, err
			}
			return postgresConnection.Connect()
		}
	case "mysql":
		{
			mysqlConnection := MySQL{
				UserName: cfg.Username,
				Password: cfg.Password,
				Host:     cfg.Host,
				Port:     cfg.Port,
				Database: cfg.Database,
				DBType:   string(cfg.DBType),
			}
			if err := mysqlConnection.IsValid(); err != nil {
				return nil, err
			}
			return mysqlConnection.Connect()
		}
	case "sqlite3":
		{
			sqlConnection := SQLite{
				Database: cfg.Database,
				DBtype:   string(cfg.DBType),
			}
			if err := sqlConnection.IsValid(); err != nil {
				return nil, err
			}
			return sqlConnection.Connect()
		}
	default:
		return nil, fmt.Errorf("invalid database type, available values: %v", database_types)
	}
}
