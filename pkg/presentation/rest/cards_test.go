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

func clearCardsTable() {
	db := mysql.DBPool()
	db.Exec("DELETE FROM cards")
	db.Exec("ALTER TABLE cards AUTO_INCREMENT = 1")
}

func TestCreateCard(t *testing.T) {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	fw, _ := writer.CreateFormField("title")
	io.Copy(fw, strings.NewReader("title1"))
	writer.Close()

	req, _ := http.NewRequest("POST", "/cards", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	response := executeRequest(req)
	checkResponseCode(t, 201, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] != "title1" {
		t.Errorf("Expected card title to be 'title1'. Got '%v'", m["title"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected card ID to be '1'. Got '%v'", m["id"])
	}

	t.Cleanup(func() {
		clearCardsTable()
	})
}

func TestUpdateCard(t *testing.T) {
	ctx := context.TODO()
	datasource.Card{}.Create(ctx, "test1", "description1")

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	fw, _ := writer.CreateFormField("title")
	io.Copy(fw, strings.NewReader("title2"))
	writer.Close()

	req, _ := http.NewRequest("PATCH", "/cards/1", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	response := executeRequest(req)
	checkResponseCode(t, 200, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != 1.0 {
		t.Errorf("Expected card ID to be '1'. Got '%v'", m["id"])
	}

	if m["title"] == "title2\n" {
		t.Errorf("Expected the title to change from 'title1' to 'title2'. Got '%v'", m["title"])
	}
}

func TestDeleteCard(t *testing.T) {
	ctx := context.TODO()
	datasource.Card{}.Create(ctx, "test1", "description1")

	req, _ := http.NewRequest("GET", "/cards/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, 200, response.Code)

	req, _ = http.NewRequest("DELETE", "/cards/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, 204, response.Code)

	req, _ = http.NewRequest("GET", "/cards/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, 422, response.Code)

	t.Cleanup(func() {
		clearCardsTable()
	})
}
