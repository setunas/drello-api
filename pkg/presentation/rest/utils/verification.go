package utils

import (
	"context"
	"drello-api/pkg/utils/firebase"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
)

func VerifyIDToken(ctx context.Context, r *http.Request) (*auth.Token, error) {
	idToken := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
	return firebase.VerifyIDToken(ctx, idToken)
}
