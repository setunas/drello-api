package controller

import (
	"drello-api/pkg/app/workspaces"
	"drello-api/pkg/infrastracture/datasource"
	"encoding/json"
	"net/http"
)

func Workspaces(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		o := workspaces.List(datasource.Workspace{})
		json.NewEncoder(w).Encode(o.Titles())
	}
}
