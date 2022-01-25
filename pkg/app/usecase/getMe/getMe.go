package getMe

import (
	"context"
	"drello-api/pkg/app/repository"
	userDM "drello-api/pkg/domain/user"
)

func Call(ctx context.Context, firebaseUID string) (*userDM.User, error) {
	user, err := (*repository.UserDS()).GetOneByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		return nil, err
	}

	return userDM.New(user.ID(), user.Username(), user.BoardID(), user.FirebaseUID()), nil
}
