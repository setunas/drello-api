package mysql

import (
	"drello-api/pkg/utils"
	"fmt"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var dbPool *sql.DB

func Open(dbUser, dbPwd, dbTCPHost, dbPort, dbName string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbTCPHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	dbPool = db
	return nil
}

func TestOpen() error {
	var (
		dbUser    = utils.MustGetenv("TEST_DB_USER")     // e.g. 'my-db-user'
		dbPwd     = utils.MustGetenv("TEST_DB_PASS")     // e.g. 'my-db-password'
		dbTCPHost = utils.MustGetenv("TEST_DB_TCP_HOST") // e.g. '127.0.0.1'
		dbPort    = utils.MustGetenv("TEST_DB_PORT")     // e.g. '3306'
		dbName    = utils.MustGetenv("TEST_DB_NAME")     // e.g. 'my-database'
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPwd, dbTCPHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
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
