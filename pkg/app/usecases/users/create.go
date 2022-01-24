package users

import (
	"context"
	"drello-api/pkg/app/repository"
	userDM "drello-api/pkg/domain/user"
)

func Create(ctx context.Context, username string, boardID int, firebaseUID string) (*userDM.User, error) {
	user, err := (*repository.UserDS()).Create(ctx, username, boardID, firebaseUID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
