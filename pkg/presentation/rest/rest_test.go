package rest

import (
	"drello-api/pkg/app/repository"
	"drello-api/pkg/infrastructure/datasource/boardsDS"
	"drello-api/pkg/infrastructure/datasource/cardsDS"
	"drello-api/pkg/infrastructure/datasource/columnsDS"
	"drello-api/pkg/infrastructure/datasource/usersDS"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/util"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	var (
		dbUser    = util.MustGetenv("TEST_DB_USER")     // e.g. 'my-db-user'
		dbPwd     = util.MustGetenv("TEST_DB_PASS")     // e.g. 'my-db-password'
		dbTCPHost = util.MustGetenv("TEST_DB_TCP_HOST") // e.g. '127.0.0.1'
		dbPort    = util.MustGetenv("TEST_DB_PORT")     // e.g. '3306'
		dbName    = util.MustGetenv("TEST_DB_NAME")     // e.g. 'my-database'
	)

	err := mysql.Open(dbUser, dbPwd, dbTCPHost, dbPort, dbName)
	if err != nil {
		log.Fatal(err)
	}

	setupDataSources()

	router = mux.NewRouter()
	setHandlers()
	os.Exit(m.Run())
}

func setupDataSources() {
	repository.SetBoardDS(boardsDS.BoardsDS{})
	repository.SetColumnDS(columnsDS.ColumnsDS{})
	repository.SetCardDS(cardsDS.CardsDS{})
	repository.SetUserDS(usersDS.UsersDS{})
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rw := httptest.NewRecorder()
	router.ServeHTTP(rw, req)
	return rw
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
