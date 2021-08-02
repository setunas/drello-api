package mysql

import (
	"drello-api/ent"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
)

var client *ent.Client

func Open() (*ent.Client, error) {
	drv, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/drello-dev")
	if err != nil {
		return nil, err
	}

	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	client = ent.NewClient(ent.Driver(drv))
	return client, nil
}

func Client() *ent.Client {
	return client
}
