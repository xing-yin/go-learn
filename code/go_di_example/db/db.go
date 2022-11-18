package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"go-learn/code/go_di_example/config"
)

func NewDatabase(config *config.Config) (*sql.DB, error) {
	return sql.Open("sqlite3", config.DatabasePath)
}
