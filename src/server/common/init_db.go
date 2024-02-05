package common

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func CreateDB() (*sql.DB, error) {
	cfg := mysql.Config{
		User:      "root",
		Passwd:    "root",
		Net:       "tcp",
		Addr:      "mysql:3306",
		DBName:    "swagger_test",
		ParseTime: true,
	}

	return sql.Open("mysql", cfg.FormatDSN())
}
