package meHandler

import (
	"drello-api/pkg/presentation/rest/restutil"
	"encoding/json"
	"net/http"
)

func get(w http.ResponseWriter, r *http.Request) error {
	user, err := restutil.AuthenticateUser(r)
	if err != nil {
		return err
	}

	json.NewEncoder(w).Encode(meResponse{
		ID:       user.ID(),
		Username: user.Username(),
		BoardID:  user.BoardID(),
	})
	return nil
}
