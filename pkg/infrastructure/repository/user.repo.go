package repository

import (
	"context"
	"drello-api/pkg/domain/user"
)

type User interface {
	GetOneByFirebaseUID(ctx context.Context, firebaseUID string) (*user.User, error)
}
