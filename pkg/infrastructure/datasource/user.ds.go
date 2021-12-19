package datasource

import (
	"context"
	"database/sql"

	userDM "drello-api/pkg/domain/user"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

type User struct{}

func (u User) GetOneByFirebaseUID(ctx context.Context, firebaseUID string) (*userDM.User, error) {
	var id int
	var username string
	var boardID int

	db := mysql.DBPool()
	row := db.QueryRow("SELECT id, username, board_id FROM users WHERE firebase_uid = ?", firebaseUID)

	switch err := row.Scan(&id, &username, &boardID); err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("not found with firebase UID %s", firebaseUID)
	case nil:
		return userDM.New(id, username, boardID, firebaseUID), nil
	default:
		return nil, err
	}
}

func (u User) Create(ctx context.Context, username string, boardID int, firebaseUID string) (*userDM.User, error) {
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
