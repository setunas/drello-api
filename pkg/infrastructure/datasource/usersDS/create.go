package usersDS

import (
	"context"

	userDM "drello-api/pkg/domain/user"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

func (u UsersDS) Create(ctx context.Context, username string, boardID int, firebaseUID string) (*userDM.User, error) {
	db := mysql.DBPool()

	result, err := db.Exec("INSERT INTO users (username, board_id, firebase_uid) VALUES (?, ?, ?)", username, boardID, firebaseUID)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return userDM.New(int(lastInsertID), username, boardID, firebaseUID), nil
}
