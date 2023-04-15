package tests

import (
	"count_on_us/config"
	db "count_on_us/internal/infrastructure/db/sqlx/connection"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose"
)

func SetupDB() *sqlx.DB {
	config := config.DatabaseConfig{
		Database: ":memory:",
		DBType:   config.SQLite3,
	}
	db, err := db.NewDatabaseFactory(&config)
	if err != nil {
		panic(err)
	}
	migrateDB(db)
	return db
}

func migrateDB(db *sqlx.DB) {
	absPath, _ := filepath.Abs("../../..")
	dir := filepath.FromSlash(fmt.Sprintf("%s/internal/infrastructure/db/sqlx/migrations", absPath))
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		panic(fmt.Errorf("directory %s not exists", dir))
	}
	if err := goose.SetDialect("sqlite3"); err != nil {
		panic(fmt.Errorf("cannot set sqlite3 dialect - error: %s", err.Error()))
	}
	if err := goose.Up(db.DB, dir); err != nil {
		panic(fmt.Errorf("cannot run migrations - error: %s", err.Error()))
	}
}
