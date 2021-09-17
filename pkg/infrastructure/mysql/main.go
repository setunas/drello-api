package mysql

import (
	"fmt"
	"log"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var dbPool *sql.DB

func Open(dbUser, dbPwd, dbTCPHost, dbPort, dbName string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/heroku_e60cb34a8aa0b6d", dbUser, dbPwd, dbTCPHost, dbPort)
	log.Println("dsn:", dsn)
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
