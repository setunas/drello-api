package main

import (
	"drello-api/pkg/app/repository"
	"drello-api/pkg/infrastructure/datasource/boardsDS"
	"drello-api/pkg/infrastructure/datasource/cardsDS"
	"drello-api/pkg/infrastructure/datasource/columnsDS"
	"drello-api/pkg/infrastructure/datasource/usersDS"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest"
	"drello-api/pkg/util"
	"drello-api/pkg/util/firebase"
	"drello-api/pkg/util/log"
)

func main() {
	setupDB()
	setupDataSources()
	firebase.InitApp()
	rest.HandleRequests()
}

func setupDB() {
	var (
		dbUser    = util.MustGetenv("DB_USER")     // e.g. 'my-db-user'
		dbPwd     = util.MustGetenv("DB_PASS")     // e.g. 'my-db-password'
		dbTCPHost = util.MustGetenv("DB_TCP_HOST") // e.g. '127.0.0.1'
		dbPort    = util.MustGetenv("DB_PORT")     // e.g. '3306'
		dbName    = util.MustGetenv("DB_NAME")     // e.g. 'my-database'
	)

	err := mysql.Open(dbUser, dbPwd, dbTCPHost, dbPort, dbName)
	if err != nil {
		log.Fatal(err)
	}
}

func setupDataSources() {
	repository.SetBoardDS(boardsDS.BoardsDS{})
	repository.SetColumnDS(columnsDS.ColumnsDS{})
	repository.SetCardDS(cardsDS.CardsDS{})
	repository.SetUserDS(usersDS.UsersDS{})
}
