package getUserByIDToken

import (
	"context"
	"drello-api/pkg/app/repository"
	"drello-api/pkg/domain/user"
)

func Call(ctx context.Context, firebaseUID string) (*user.User, error) {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
