package columnsHandler

// import (
// 	"bytes"
// 	"context"
// 	"drello-api/pkg/app/repository"
// 	"drello-api/pkg/infrastructure/mysql"
// 	"encoding/json"
// 	"io"
// 	"mime/multipart"
// 	"net/http"
// 	"strings"
// 	"testing"
// )

// func clearColumnsTable() {
// 	db := mysql.DBPool()
// 	db.Exec("DELETE FROM columns")
// 	db.Exec("ALTER TABLE columns AUTO_INCREMENT = 1")
// }

// func TestCreateColumn(t *testing.T) {
// 	ctx := context.TODO()
// 	(*repository.BoardDS()).Create(ctx, "test1")

// 	var body bytes.Buffer
// 	writer := multipart.NewWriter(&body)

// 	fw, _ := writer.CreateFormField("title")
// 	io.Copy(fw, strings.NewReader("title1"))

// 	fw, _ = writer.CreateFormField("boardId")
// 	io.Copy(fw, strings.NewReader("1"))

// 	writer.Close()

// 	req, _ := http.NewRequest("POST", "/columns", &body)
// 	req.Header.Set("Content-Type", writer.FormDataContentType())

// 	response := executeRequest(req)
// 	checkResponseCode(t, 201, response.Code)

// 	var m map[string]interface{}
// 	json.Unmarshal(response.Body.Bytes(), &m)

// 	if m["title"] != "title1" {
// 		t.Errorf("Expected the title to be 'title1'. Got '%v'", m["title"])
// 	}

// 	if m["id"] != 1.0 {
// 		t.Errorf("Expected column ID to be '1'. Got '%v'", m["id"])
// 	}

// 	if m["boardId"] != 1.0 {
// 		t.Errorf("Expected boardId to be '1'. Got '%v'", m["boardId"])
// 	}

// 	t.Cleanup(func() {
// 		clearBoardsTable()
// 		clearColumnsTable()
// 	})
// }

// func TestUpdateColumn(t *testing.T) {
// 	ctx := context.TODO()
// 	(*repository.BoardDS()).Create(ctx, "test1")
// 	(*repository.ColumnDS()).Create(ctx, "test1", 1.0, 1)

// 	var body bytes.Buffer
// 	writer := multipart.NewWriter(&body)

// 	fw, _ := writer.CreateFormField("title")
// 	io.Copy(fw, strings.NewReader("title2"))

// 	fw, _ = writer.CreateFormField("boardId")
// 	io.Copy(fw, strings.NewReader("2"))

// 	writer.Close()

// 	req, _ := http.NewRequest("PATCH", "/columns/1", &body)
// 	req.Header.Set("Content-Type", writer.FormDataContentType())

// 	response := executeRequest(req)
// 	checkResponseCode(t, 200, response.Code)

// 	var m map[string]interface{}
// 	json.Unmarshal(response.Body.Bytes(), &m)

// 	if m["id"] != 1.0 {
// 		t.Errorf("Expected column ID to be '1'. Got '%v'", m["id"])
// 	}

// 	if m["title"] == "title2\n" {
// 		t.Errorf("Expected the title to change from 'title1' to 'title2'. Got '%v'", m["title"])
// 	}

// 	if m["boardId"] != 2.0 {
// 		t.Errorf("Expected boardId to be '2'. Got '%v'", m["boardId"])
// 	}

// 	t.Cleanup(func() {
// 		clearBoardsTable()
// 		clearColumnsTable()
// 	})
// }

// func TestDeleteColumn(t *testing.T) {
// 	ctx := context.TODO()
// 	(*repository.BoardDS()).Create(ctx, "test1")
// 	(*repository.ColumnDS()).Create(ctx, "test1", 1.0, 1)

// 	req, _ := http.NewRequest("DELETE", "/columns/1", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, 204, response.Code)

// 	t.Cleanup(func() {
// 		clearBoardsTable()
// 		clearColumnsTable()
// 	})
// }
