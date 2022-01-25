package rest

import (
	"drello-api/pkg/app/usecase/signup"
	"encoding/json"
	"net/http"
)

type signupResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	BoardID  int    `json:"boardId"`
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodOptions:
		return

	case http.MethodPost:
		token, err := verifyIDToken(r.Context(), r)
		if err != nil {
			handleClientError(w, err, 401, "Invalid token")
			return
		}

		var body struct {
			Username string
			Title    string
		}
		json.NewDecoder(r.Body).Decode(&body)

		user, err := signup.Signup(r.Context(), body.Username, token.UID, body.Title)
		if err != nil {
			handleClientError(w, err, 422, "An error occured while creating a user")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(signupResponse{ID: user.ID(), Username: user.Username(), BoardID: user.BoardID()})
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}
