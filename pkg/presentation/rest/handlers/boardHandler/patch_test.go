package boardHandler

import (
	"bytes"
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/presentation/rest/resttest"
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
	(*repository.UserDS()).Create(ctx, "user1", 1, "UID-1")

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	fw, _ := writer.CreateFormField("title")
	io.Copy(fw, strings.NewReader("title2"))
	writer.Close()

	req, _ := http.NewRequest("PATCH", "/boards/1", &body)
	req.Header.Set("Authorization", "Bearer UID-1")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	response := resttest.ExecuteRequest(req)
	resttest.CheckResponseCode(t, 200, response)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != 1.0 {
		t.Errorf("Expected board ID to be '1'. Got '%v'", m["id"])
	}

	if m["title"] == "title2\n" {
		t.Errorf("Expected board title to change from 'title1' to 'title2'. Got '%v'", m["title"])
	}

	t.Cleanup(func() {
		resttest.CleanupAllTable()
	})
}
