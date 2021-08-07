package rest

import (
	"drello-api/pkg/app/workspaces"
	"drello-api/pkg/constants"
	"drello-api/pkg/infrastracture/datasource"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func Workspaces(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		output, err := workspaces.List(r.Context(), datasource.Workspace{})
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		var wolist []*resWorkspace
		for _, wo := range output.Workspaces {
			wolist = append(wolist, &resWorkspace{
				ID:    wo.ID(),
				Title: wo.Title(),
			})
		}

		json.NewEncoder(w).Encode(wolist)

	case http.MethodPost:
		output, err := workspaces.Create(r.Context(), datasource.Workspace{}, &workspaces.CreateInput{Title: r.FormValue("title")})
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resWorkspace{ID: output.Workspace.ID(), Title: output.Workspace.Title()})

	case http.MethodPatch:
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, constants.Workspaces))
		if err != nil {
			handleClientError(w, err, 400, "Invalid ID.")
			return
		}

		output, err := workspaces.Update(r.Context(), datasource.Workspace{}, &workspaces.UpdateInput{ID: id, Title: r.FormValue("title")})
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(resWorkspace{ID: output.Workspace.ID(), Title: output.Workspace.Title()})

	case http.MethodDelete:
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, constants.Workspaces))
		if err != nil {
			handleClientError(w, err, 400, "Invalid ID.")
		}

		err = workspaces.Delete(r.Context(), datasource.Workspace{}, workspaces.NewDeleteInput(id))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

type resWorkspace struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
