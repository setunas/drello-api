package rest

import (
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	var (
		dbUser    = utils.MustGetenv("TEST_DB_USER")     // e.g. 'my-db-user'
		dbPwd     = utils.MustGetenv("TEST_DB_PASS")     // e.g. 'my-db-password'
		dbTCPHost = utils.MustGetenv("TEST_DB_TCP_HOST") // e.g. '127.0.0.1'
		dbPort    = utils.MustGetenv("TEST_DB_PORT")     // e.g. '3306'
		dbName    = utils.MustGetenv("TEST_DB_NAME")     // e.g. 'my-database'
	)

	err := mysql.Open(dbUser, dbPwd, dbTCPHost, dbPort, dbName)
	if err != nil {
		log.Println(err)
	}

	router = mux.NewRouter()
	setHandlers()
	os.Exit(m.Run())
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
