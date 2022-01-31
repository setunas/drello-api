package boardHandler

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/presentation/rest/resttest"
	"encoding/json"
	"net/http"
	"testing"
)

func TestGetBoardRequest(t *testing.T) {
	ctx := context.TODO()
	(*repository.BoardDS()).Create(ctx, "test1")
	(*repository.BoardDS()).Create(ctx, "test2")
	(*repository.ColumnDS()).Create(ctx, "test1", 1.0, 1)
	(*repository.ColumnDS()).Create(ctx, "test2", 2.0, 2)
	(*repository.CardDS()).Create(ctx, "test1", "desc1", 1.0, 1)
	(*repository.CardDS()).Create(ctx, "test2", "desc2", 2.0, 2)
	(*repository.UserDS()).Create(ctx, "user1", 1, "UID-1")

	req, _ := http.NewRequest("GET", "/boards/3", nil)
	req.Header.Set("Authorization", "Bearer UID-1")
	response := resttest.ExecuteRequest(req)
	resttest.CheckResponseCode(t, 500, response)

	req, _ = http.NewRequest("GET", "/boards/1", nil)
	req.Header.Set("Authorization", "Bearer UID-2")
	response = resttest.ExecuteRequest(req)
	resttest.CheckResponseCode(t, 401, response)

	req, _ = http.NewRequest("GET", "/boards/1", nil)
	req.Header.Set("Authorization", "Bearer UID-1")
	response = resttest.ExecuteRequest(req)
	resttest.CheckResponseCode(t, 200, response)

	var wr boardResponse
	json.Unmarshal(response.Body.Bytes(), &wr)

	expectedBody := `{"id":1,"title":"test1","columns":[{"id":1,"title":"test1","position":1,"boardId":1}],"cards":[{"id":1,"title":"test1","description":"desc1","position":1,"columnId":1}]}` + "\n"
	if body := response.Body.String(); body != expectedBody {
		t.Errorf("Expected %s. Got %s", expectedBody, body)
	}

	t.Cleanup(func() {
		resttest.CleanupAllTable()
	})
}
