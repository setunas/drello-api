package resttest

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/infrastructure/datasource/boardsDS"
	"drello-api/pkg/infrastructure/datasource/cardsDS"
	"drello-api/pkg/infrastructure/datasource/columnsDS"
	"drello-api/pkg/infrastructure/datasource/usersDS"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/util"
	"drello-api/pkg/util/firebase"
	"drello-api/pkg/util/log"
	"runtime"

	"net/http"
	"net/http/httptest"
	"testing"

	"firebase.google.com/go/v4/auth"
	"github.com/gorilla/mux"
)

var router *mux.Router

func Router() *mux.Router {
	return router
}

func InitTest() {
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

	firebase.SetVerifyIDToken(verifyIDTokenImpl)
	setupDataSources()
	router = mux.NewRouter()
}

func verifyIDTokenImpl(ctx context.Context, idToken string) (*auth.Token, error) {
	token := &auth.Token{UID: idToken}
	return token, nil
}

func setupDataSources() {
	repository.SetBoardDS(boardsDS.BoardsDS{})
	repository.SetColumnDS(columnsDS.ColumnsDS{})
	repository.SetCardDS(cardsDS.CardsDS{})
	repository.SetUserDS(usersDS.UsersDS{})
}

func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rw := httptest.NewRecorder()
	router.ServeHTTP(rw, req)
	return rw
}

func CheckResponseCode(t *testing.T, expectedCode int, actualRes *httptest.ResponseRecorder) {
	_, filename, line, _ := runtime.Caller(1)
	if expectedCode != actualRes.Code {
		t.Errorf("%s:%d Expected response code %d. Got %d. %s\n", filename, line, expectedCode, actualRes.Code, actualRes.Body)
	}
}

func CleanupAllTable() {
	db := mysql.DBPool()

	db.Exec("DELETE FROM cards")
	db.Exec("ALTER TABLE cards AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM columns")
	db.Exec("ALTER TABLE columns AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM users")
	db.Exec("ALTER TABLE users AUTO_INCREMENT = 1")

	db.Exec("DELETE FROM boards")
	db.Exec("ALTER TABLE boards AUTO_INCREMENT = 1")
}
