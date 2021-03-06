package restutil

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const requestID = "requestID"

func AppendReqIDToCtx(r *http.Request) *http.Request {
	id := uuid.New()
	ctx := context.WithValue(r.Context(), requestID, id.String())
	return r.WithContext(ctx)
}

func RetrieveReqID(ctx context.Context) string {
	reqIDRaw := ctx.Value(requestID)
	reqID, ok := reqIDRaw.(string)
	if !ok {
		reqID = "undefined"
	}

	return reqID
}
