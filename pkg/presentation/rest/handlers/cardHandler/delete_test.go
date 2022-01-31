package cardHandler

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/presentation/rest/resttest"
	"net/http"
	"testing"
)

func TestDeleteCard(t *testing.T) {
	ctx := context.TODO()
	(*repository.BoardDS()).Create(ctx, "test1")
	(*repository.ColumnDS()).Create(ctx, "test1", 1.0, 1)
	(*repository.CardDS()).Create(ctx, "test1", "description1", 1.0, 1)
	(*repository.UserDS()).Create(ctx, "user1", 1, "UID-1")

	req, _ := http.NewRequest("DELETE", "/cards/1", nil)
	req.Header.Set("Authorization", "Bearer UID-1")
	response := resttest.ExecuteRequest(req)
	resttest.CheckResponseCode(t, 204, response)
}
