package controller

import (
	"drello-api/pkg/app/workspaces"
	"drello-api/pkg/infrastracture/datasource"
	"encoding/json"
	"fmt"
	"net/http"
)

func Workspaces(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		o, err := workspaces.List(r.Context(), datasource.Workspace{})
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(o.Titles())
	}
}
