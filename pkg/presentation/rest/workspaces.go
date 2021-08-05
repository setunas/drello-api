package rest

import (
	"drello-api/pkg/app/workspaces"
	"drello-api/pkg/infrastracture/datasource"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Workspaces(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		output, err := workspaces.List(r.Context(), datasource.Workspace{})
		if err != nil {
			fmt.Println(err)
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
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(resWorkspace{ID: output.Workspace.ID(), Title: output.Workspace.Title()})

	case http.MethodPatch:
		id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/workspaces/"))
		if err != nil {
			fmt.Println(err)
		}

		output, err := workspaces.Update(r.Context(), datasource.Workspace{}, &workspaces.UpdateInput{ID: id, Title: r.FormValue("title")})
		if err != nil {
			fmt.Println(err)
		}

		json.NewEncoder(w).Encode(resWorkspace{ID: output.Workspace.ID(), Title: output.Workspace.Title()})
	}
}

type resWorkspace struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
