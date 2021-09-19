package rest

import (
	"bytes"
	"context"
	"drello-api/pkg/infrastructure/datasource"
	"drello-api/pkg/infrastructure/mysql"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
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
	os.Exit(m.Run())
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
		clearWorkspacesTable()
	})
}

func TestGetWorkspace(t *testing.T) {
	ctx := context.TODO()
	datasource.Workspace{}.Create(ctx, "test1")

	req, _ := http.NewRequest("GET", "/workspaces/2", nil)
	response := executeRequest(req)
	checkResponseCode(t, 422, response.Code)

	req, _ = http.NewRequest("GET", "/workspaces/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, 200, response.Code)

	var wr workspaceResponse
	json.Unmarshal(response.Body.Bytes(), &wr)

	expectedBody := `{"id":1,"title":"test1"}` + "\n"
	if body := response.Body.String(); body != expectedBody {
		t.Errorf("Expected %s. Got %s", expectedBody, body)
	}

	t.Cleanup(func() {
		clearWorkspacesTable()
	})
}

func TestCreateWorkspace(t *testing.T) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	fw, _ := writer.CreateFormField("title")
	io.Copy(fw, strings.NewReader("title1"))
	writer.Close()

	req, _ := http.NewRequest("POST", "/workspaces", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	response := executeRequest(req)
	checkResponseCode(t, 201, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] != "title1" {
		t.Errorf("Expected workspace title to be 'title1'. Got '%v'", m["title"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected workspace ID to be '1'. Got '%v'", m["id"])
	}

	t.Cleanup(func() {
		clearWorkspacesTable()
	})
}

func TestUpdateWorkspace(t *testing.T) {
	ctx := context.TODO()
	datasource.Workspace{}.Create(ctx, "test1")

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	fw, _ := writer.CreateFormField("title")
	io.Copy(fw, strings.NewReader("title2"))
	writer.Close()

	req, _ := http.NewRequest("PATCH", "/workspaces/1", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	response := executeRequest(req)
	checkResponseCode(t, 200, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != 1.0 {
		t.Errorf("Expected workspace ID to be '1'. Got '%v'", m["id"])
	}

	if m["title"] == "title2\n" {
		t.Errorf("Expected the title to change from 'title1' to 'title2'. Got '%v'", m["title"])
	}
}

func TestDeleteWorkspace(t *testing.T) {
	ctx := context.TODO()
	datasource.Workspace{}.Create(ctx, "test1")

	req, _ := http.NewRequest("GET", "/workspaces/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, 200, response.Code)

	req, _ = http.NewRequest("DELETE", "/workspaces/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, 204, response.Code)

	req, _ = http.NewRequest("GET", "/workspaces/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, 422, response.Code)

	t.Cleanup(func() {
		clearWorkspacesTable()
	})
}
