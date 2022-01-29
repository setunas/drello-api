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

func RetrieveReqID(r *http.Request) string {
	reqIDRaw := r.Context().Value(requestID)
	reqID, ok := reqIDRaw.(string)
	if !ok {
		reqID = "undefined"
	}

	return reqID
}
