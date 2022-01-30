package firebase

import (
	"context"
	"drello-api/pkg/util/apperr"
	"drello-api/pkg/util/log"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

var firebaseApp *firebase.App

func InitApp() {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal("Error initializing firebase app: %v\n", err)
	}

	firebaseApp = app
	verifyIDToken = verifyIDTokenImpl
}

type verifyIDTokenType func(ctx context.Context, idToken string) (*auth.Token, error)

var verifyIDToken verifyIDTokenType

func VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	return verifyIDToken(ctx, idToken)
}

func SetVerifyIDToken(fn verifyIDTokenType) {
	verifyIDToken = fn
}

func verifyIDTokenImpl(ctx context.Context, idToken string) (*auth.Token, error) {
	client, err := firebaseApp.Auth(ctx)
	if err != nil {
		return nil, apperr.NewAppError([]apperr.Tag{apperr.FailedAuthorization}, fmt.Sprintf("error getting Auth client: %v", err), err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, apperr.NewAppError([]apperr.Tag{apperr.FailedAuthorization}, fmt.Sprintf("error verifying ID token: %v", err), err)
	}

	return token, nil
}
