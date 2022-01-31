package columnsHandler

import (
	"bytes"
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/presentation/rest/resttest"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateColumn(t *testing.T) {
	ctx := context.TODO()
	(*repository.BoardDS()).Create(ctx, "test1")
	(*repository.UserDS()).Create(ctx, "user1", 1, "UID-1")

	jsonStr := []byte(`{
		"title":"title1",
		"BoardID":1
	}`)
	req, _ := http.NewRequest("POST", "/columns", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer UID-1")
	response := resttest.ExecuteRequest(req)
	resttest.CheckResponseCode(t, 201, response)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] != "title1" {
		t.Errorf("Expected the title to be 'title1'. Got '%v'", m["title"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected column ID to be '1'. Got '%v'", m["id"])
	}

	if m["boardId"] != 1.0 {
		t.Errorf("Expected boardId to be '1'. Got '%v'", m["boardId"])
	}

	t.Cleanup(func() {
		resttest.CleanupAllTable()
	})
}
