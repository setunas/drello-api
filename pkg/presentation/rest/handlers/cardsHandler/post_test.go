package cardsHandler

import (
	"bytes"
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/presentation/rest/resttest"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateCard(t *testing.T) {
	ctx := context.TODO()
	(*repository.BoardDS()).Create(ctx, "test1")
	(*repository.ColumnDS()).Create(ctx, "test1", 1.0, 1)
	(*repository.UserDS()).Create(ctx, "user1", 1, "UID-1")

	jsonStr := []byte(`{
		"title":"title1",
		"ColumnID":1
	}`)
	req, _ := http.NewRequest("POST", "/cards", bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer UID-1")
	req.Header.Set("Content-Type", "application/json")

	response := resttest.ExecuteRequest(req)
	resttest.CheckResponseCode(t, 201, response)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["title"] != "title1" {
		t.Errorf("Expected card title to be 'title1'. Got '%v'", m["title"])
	}

	if m["id"] != 1.0 {
		t.Errorf("Expected card ID to be '1'. Got '%v'", m["id"])
	}

	if m["columnId"] != 1.0 {
		t.Errorf("Expected columnId to be '1'. Got '%v'", m["columnId"])
	}

	t.Cleanup(func() {
		resttest.CleanupAllTable()
	})
}
