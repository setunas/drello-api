package util

import (
	"context"
	"drello-api/pkg/app/usecase/getUserByIDToken"
	"drello-api/pkg/domain/user"
	"drello-api/pkg/util/firebase"
	"drello-api/pkg/util/myerr"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
)

func VerifyIDToken(ctx context.Context, r *http.Request) (*auth.Token, error) {
	idToken := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)
	return firebase.VerifyIDToken(ctx, idToken)
}

func AuthenticateUser(r *http.Request) (*user.User, error) {
	token, err := VerifyIDToken(r.Context(), r)
	if err != nil {
		return nil, myerr.NewHTTPError(401, "Invalid token", err)
	}
	return getUserByIDToken.Call(r.Context(), token.UID)
}
