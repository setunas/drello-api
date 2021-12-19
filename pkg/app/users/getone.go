package users

import (
	"context"
	userDM "drello-api/pkg/domain/user"
	"drello-api/pkg/infrastructure/repository"
)

func GetOne(ctx context.Context, userRepo repository.User, input *GetOneInput) (*GetOneOutput, error) {
	user, err := userRepo.GetOneByFirebaseUID(ctx, input.firebaseUID)
	if err != nil {
		return nil, err
	}

	return &GetOneOutput{User: userDM.New(user.ID(), user.Username(), user.BoardID(), user.FirebaseUID())}, nil
}

type GetOneInput struct {
	firebaseUID string
}

func NewGetOneInput(firebaseUID string) *GetOneInput {
	return &GetOneInput{firebaseUID: firebaseUID}
}

type GetOneOutput struct {
	User *userDM.User
}
