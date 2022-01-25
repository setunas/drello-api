package usersDS

import (
	"context"
	"database/sql"

	userDM "drello-api/pkg/domain/user"
	"drello-api/pkg/infrastructure/mysql"
	"fmt"
)

func (u UsersDS) GetOneByFirebaseUID(ctx context.Context, firebaseUID string) (*userDM.User, error) {
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
