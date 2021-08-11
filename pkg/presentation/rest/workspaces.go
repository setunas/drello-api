package rest

import (
	"drello-api/pkg/app/workspaces"
	"drello-api/pkg/infrastracture/datasource"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type workspaceResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func workspacesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		output, err := workspaces.GetAll(r.Context(), datasource.Workspace{})
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		var wolist []*workspaceResponse
		for _, wo := range output.Workspaces {
			wolist = append(wolist, &workspaceResponse{
				ID:    wo.ID(),
				Title: wo.Title(),
			})
		}

		json.NewEncoder(w).Encode(wolist)

	case http.MethodPost:
		output, err := workspaces.Create(r.Context(), datasource.Workspace{}, workspaces.NewCreateInput(r.FormValue("title")))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(workspaceResponse{ID: output.Workspace.ID(), Title: output.Workspace.Title()})
	}

	handleClientError(w, nil, 404, "Invalid path")
}

func workspaceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		handleClientError(w, err, 400, "Invalid ID.")
	}

	switch r.Method {
	case http.MethodGet:
		output, err := workspaces.GetOne(r.Context(), datasource.Workspace{}, workspaces.NewGetOneInput(id))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(&workspaceResponse{
			ID:    output.Workspace.ID(),
			Title: output.Workspace.Title(),
		})

	case http.MethodPatch:
		output, err := workspaces.Update(r.Context(), datasource.Workspace{}, workspaces.NewUpdateInput(id, r.FormValue("title")))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		json.NewEncoder(w).Encode(workspaceResponse{ID: output.Workspace.ID(), Title: output.Workspace.Title()})

	case http.MethodDelete:
		err = workspaces.Delete(r.Context(), datasource.Workspace{}, workspaces.NewDeleteInput(id))
		if err != nil {
			handleClientError(w, err, 422, "An error occured during the prosess")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}

	handleClientError(w, nil, 404, "Invalid path")
}
