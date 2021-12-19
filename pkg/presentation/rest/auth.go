package rest

import (
	"context"
	"drello-api/pkg/utils/firebase"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
)

func verifyIDToken(ctx context.Context, r *http.Request) *auth.Token {
	idToken := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
	if idToken == "" {
		println("No ID Token founded")
	}

	token := firebase.VerifyIDToken(ctx, idToken)
	log.Printf("Verified ID token: %v\n", token)

	return token
}
