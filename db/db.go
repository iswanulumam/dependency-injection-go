package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewSql(mysqlHost, mysqlUser, mysqlPassword string) (*sql.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:3306)/following", mysqlUser, mysqlPassword, mysqlHost)
	sqlDB, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	return sqlDB, nil
}
