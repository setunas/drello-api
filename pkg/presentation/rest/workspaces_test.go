package rest

import (
	"context"
	"drello-api/pkg/infrastructure/datasource"
	"drello-api/pkg/infrastructure/mysql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(m *testing.M) {
	err := mysql.Open("root", "password", "127.0.0.1", "4306", "drello-test")
	if err != nil {
		log.Println(err)
	}

	router = mux.NewRouter()
	setHandlers()

	clearWorkspacesTable()
	code := m.Run()
	clearWorkspacesTable()
	os.Exit(code)
}

func clearWorkspacesTable() {
	db := mysql.DBPool()
	db.Exec("DELETE FROM workspaces")
	db.Exec("ALTER TABLE workspaces AUTO_INCREMENT = 1")
}

func TestGetWorkspaces(t *testing.T) {
	ctx := context.TODO()
	datasource.Workspace{}.Create(ctx, "test1")
	datasource.Workspace{}.Create(ctx, "test2")
	datasource.Workspace{}.Create(ctx, "test3")

	req, _ := http.NewRequest("GET", "/workspaces", nil)
	response := executeRequest(req)

	checkResponseCode(t, 200, response.Code)

	var wr []workspaceResponse
	json.Unmarshal(response.Body.Bytes(), &wr)

	if len(wr) != 3 {
		t.Errorf("Expected 3. Got %d", len(wr))
	}

	expectedBody := `[{"id":1,"title":"test1"},{"id":2,"title":"test2"},{"id":3,"title":"test3"}]` + "\n"
	if body := response.Body.String(); body != expectedBody {
		t.Errorf("Expected %s. Got %s", expectedBody, body)
	}

	t.Cleanup(func() {

	})
}

// func TestDeleteProduct(t *testing.T) {
// 	addProducts(1)

// 	req, _ := http.NewRequest("GET", "/product/1", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	req, _ = http.NewRequest("DELETE", "/product/1", nil)
// 	response = executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	req, _ = http.NewRequest("GET", "/product/1", nil)
// 	response = executeRequest(req)
// 	checkResponseCode(t, http.StatusNotFound, response.Code)
// }
