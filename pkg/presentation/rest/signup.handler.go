package rest

import (
	"drello-api/pkg/app/signup"
	"drello-api/pkg/infrastructure/datasource"
	"encoding/json"
	"net/http"
)

type signupResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	BoardID  int    `json:"boardId"`
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodPost:
		token, err := verifyIDToken(ctx, r)
		if err != nil {
			handleClientError(w, err, 401, "Invalid token")
			return
		}

		var body struct {
			Username string
			Title    string
		}
		json.NewDecoder(r.Body).Decode(&body)

		signupOutput, err := signup.Signup(r.Context(), datasource.User{}, datasource.Board{}, datasource.Column{}, signup.NewSignupInput(body.Username, token.UID, body.Title))
		if err != nil {
			handleClientError(w, err, 422, "An error occured while creating a user")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(signupResponse{ID: signupOutput.User.ID(), Username: signupOutput.User.Username(), BoardID: signupOutput.User.BoardID()})
		return
	}

	handleClientError(w, nil, 404, "Invalid method")
}
