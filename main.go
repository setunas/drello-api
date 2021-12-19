package main

import (
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest"
	"drello-api/pkg/utils"
	"drello-api/pkg/utils/firebase"
	"log"
)

func main() {
	var (
		dbUser    = utils.MustGetenv("DB_USER")     // e.g. 'my-db-user'
		dbPwd     = utils.MustGetenv("DB_PASS")     // e.g. 'my-db-password'
		dbTCPHost = utils.MustGetenv("DB_TCP_HOST") // e.g. '127.0.0.1'
		dbPort    = utils.MustGetenv("DB_PORT")     // e.g. '3306'
		dbName    = utils.MustGetenv("DB_NAME")     // e.g. 'my-database'
	)

	err := mysql.Open(dbUser, dbPwd, dbTCPHost, dbPort, dbName)
	if err != nil {
		log.Println(err)
	}

	firebase.InitApp()
	rest.HandleRequests()
}
