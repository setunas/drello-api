package mysql

import (
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var dbPool *sql.DB

func Open() error {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:4306)/drello-dev")
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	dbPool = db
	return nil
}

func DBPool() *sql.DB {
	return dbPool
}
