package boardHandler

import (
	"bytes"
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest/test"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"testing"
)

func TestPatchBoardRequest(t *testing.T) {
	ctx := context.TODO()
	(*repository.BoardDS()).Create(ctx, "test1")

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	fw, _ := writer.CreateFormField("title")
	io.Copy(fw, strings.NewReader("title2"))
	writer.Close()

	req, _ := http.NewRequest("PATCH", "/boards/1", &body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	response := test.ExecuteRequest(req)
	test.CheckResponseCode(t, 200, response)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != 1.0 {
		t.Errorf("Expected board ID to be '1'. Got '%v'", m["id"])
	}

	if m["title"] == "title2\n" {
		t.Errorf("Expected board title to change from 'title1' to 'title2'. Got '%v'", m["title"])
	}

	t.Cleanup(func() {
		db := mysql.DBPool()
		db.Exec("DELETE FROM boards")
		db.Exec("ALTER TABLE boards AUTO_INCREMENT = 1")
	})
}
