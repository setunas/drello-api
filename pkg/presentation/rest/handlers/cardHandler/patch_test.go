package cardHandler

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

func TestUpdateCard(t *testing.T) {
	ctx := context.TODO()
	(*repository.BoardDS()).Create(ctx, "test1")
	(*repository.ColumnDS()).Create(ctx, "test1", 1.0, 1)
	(*repository.ColumnDS()).Create(ctx, "test2", 2.0, 1)
	(*repository.CardDS()).Create(ctx, "test1", "desc1", 1.0, 1)
	(*repository.UserDS()).Create(ctx, "user1", 1, "UID-1")

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	fw, _ := writer.CreateFormField("title")
	io.Copy(fw, strings.NewReader("title2"))

	fw, _ = writer.CreateFormField("description")
	io.Copy(fw, strings.NewReader("desc2"))

	fw, _ = writer.CreateFormField("columnId")
	io.Copy(fw, strings.NewReader("2"))

	writer.Close()

	jsonStr := []byte(`{
		"title":"title2",
		"ColumnID":2
	}`)
	req, _ := http.NewRequest("PATCH", "/cards/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer UID-1")
	req.Header.Set("Content-Type", writer.FormDataContentType())

	response := resttest.ExecuteRequest(req)
	resttest.CheckResponseCode(t, 200, response)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != 1.0 {
		t.Errorf("Expected card ID to be '1'. Got '%v'", m["id"])
	}

	if m["title"] == "title2\n" {
		t.Errorf("Expected the title to change from 'title1' to 'title2'. Got '%v'", m["title"])
	}

	if m["columnId"] != 2.0 {
		t.Errorf("Expected columnId to be '2'. Got '%v'", m["columnId"])
	}
	t.Cleanup(func() {
		resttest.CleanupAllTable()
	})
}
