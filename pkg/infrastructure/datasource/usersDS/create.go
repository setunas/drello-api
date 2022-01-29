package usersDS

import (
	"context"

	userDM "drello-api/pkg/domain/user"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/util/log"
	"fmt"
)

func (u UsersDS) Create(ctx context.Context, username string, boardID int, firebaseUID string) (*userDM.User, error) {
	db := mysql.DBPool()

	query := "INSERT INTO users (username, board_id, firebase_uid) VALUES (?, ?, ?)"
	log.Info("SQL").Add("SQL", query).Add("username", username).Add("boardID", boardID).
		Add("firebaseUID", firebaseUID).Write()
	result, err := db.Exec(query, username, boardID, firebaseUID)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return userDM.New(int(lastInsertID), username, boardID, firebaseUID), nil
}
