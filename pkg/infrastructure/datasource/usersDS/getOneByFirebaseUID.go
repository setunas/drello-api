package usersDS

import (
	"context"
	"database/sql"

	userDM "drello-api/pkg/domain/user"
	"drello-api/pkg/infrastructure/mysql"
	"drello-api/pkg/presentation/rest/restutil"
	"drello-api/pkg/util/apperr"
	"drello-api/pkg/util/log"
	"fmt"
)

func (u UsersDS) GetOneByFirebaseUID(ctx context.Context, firebaseUID string) (*userDM.User, error) {
	var id int
	var username string
	var boardID int

	db := mysql.DBPool()
	query := "SELECT id, username, board_id FROM users WHERE firebase_uid = ?"
	log.Info("SQL").Add("RequestID", restutil.RetrieveReqID(ctx)).Add("SQL", query).Add("firebaseUID", firebaseUID).Write()
	row := db.QueryRow(query, firebaseUID)

	switch err := row.Scan(&id, &username, &boardID); err {
	case sql.ErrNoRows:
		return nil, apperr.NewAppError([]apperr.Tag{apperr.RecordNotFound}, fmt.Sprintf("not found with firebase UID %s", firebaseUID), nil)
	case nil:
		return userDM.New(id, username, boardID, firebaseUID), nil
	default:
		return nil, err
	}
}
