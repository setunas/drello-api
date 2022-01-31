package columnHandler

import (
	"bytes"
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/presentation/rest/resttest"
	"encoding/json"
	"net/http"
	"testing"
)

func TestUpdateColumn(t *testing.T) {
	ctx := context.TODO()
	(*repository.BoardDS()).Create(ctx, "test1")
	(*repository.ColumnDS()).Create(ctx, "test1", 1.0, 1)
	(*repository.UserDS()).Create(ctx, "user1", 1, "UID-1")

	jsonStr := []byte(`{
		"title":"title2",
		"BoardID":1
	}`)
	req, _ := http.NewRequest("PATCH", "/columns/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer UID-1")
	response := resttest.ExecuteRequest(req)
	resttest.CheckResponseCode(t, 200, response)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != 1.0 {
		t.Errorf("Expected column ID to be '1'. Got '%v'", m["id"])
	}

	if m["title"] == "title2\n" {
		t.Errorf("Expected the title to change from 'title1' to 'title2'. Got '%v'", m["title"])
	}

	if m["boardId"] != 1.0 {
		t.Errorf("Expected boardId to be '2'. Got '%v'", m["boardId"])
	}

	t.Cleanup(func() {
		resttest.CleanupAllTable()
	})
}
