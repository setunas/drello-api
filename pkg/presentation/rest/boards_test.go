package rest

import (
	"bytes"
	"context"
	"drello-api/pkg/infrastructure/datasource"
	"drello-api/pkg/infrastructure/mysql"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"testing"
)

func clearBoardsTable() {
	db := mysql.DBPool()
	db.Exec("DELETE FROM boards")
	db.Exec("ALTER TABLE boards AUTO_INCREMENT = 1")
}

func TestGetBoard(t *testing.T) {
	ctx := context.TODO()
	datasource.Board{}.Create(ctx, "test1")
	datasource.Board{}.Create(ctx, "test2")
	datasource.Column{}.Create(ctx, "test1", 1)
	datasource.Column{}.Create(ctx, "test2", 2)
	datasource.Card{}.Create(ctx, "test1", "desc1", 1)
	datasource.Card{}.Create(ctx, "test2", "desc2", 2)

	req, _ := http.NewRequest("GET", "/boards/3", nil)
	response := executeRequest(req)
	checkResponseCode(t, 422, response.Code)

	req, _ = http.NewRequest("GET", "/boards/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, 200, response.Code)

	var wr boardResponse
	json.Unmarshal(response.Body.Bytes(), &wr)

	expectedBody := `{"id":1,"title":"test1","columns":[{"id":1,"title":"test1","boardId":1}],"cards":[{"id":1,"title":"test1","description":"desc1","columnId":1}]}` + "\n"
	if body := response.Body.String(); body != expectedBody {
		t.Errorf("Expected %s. Got %s", expectedBody, body)
	}

	t.Cleanup(func() {
		clearBoardsTable()
		clearColumnsTable()
		clearCardsTable()
	})
}

func TestUpdateBoard(t *testing.T) {
	ctx := context.TODO()
	datasource.Board{}.Create(ctx, "test1")

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	fw, _ := writer.CreateFormField("title")
	io.Copy(fw, strings.NewReader("title2"))
	writer.Close()

	req, _ := http.NewRequest("PATCH", "/boards/1", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	response := executeRequest(req)
	checkResponseCode(t, 200, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != 1.0 {
		t.Errorf("Expected board ID to be '1'. Got '%v'", m["id"])
	}

	if m["title"] == "title2\n" {
		t.Errorf("Expected board title to change from 'title1' to 'title2'. Got '%v'", m["title"])
	}

	t.Cleanup(func() {
		clearBoardsTable()
	})
}
