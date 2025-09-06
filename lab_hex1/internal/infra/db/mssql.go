package db

import (
	"database/sql"
	"fmt"
	"hex1/internal/config"

	_ "github.com/denisenkom/go-mssqldb"
)

func NewSQLServer(cfg config.SQLServerConfig) (*sql.DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Port, cfg.Database)
	return sql.Open("sqlserver", connString)
}
