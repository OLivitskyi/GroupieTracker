package handlers

import (
	"GroupieTracker/data"
	"encoding/json"
	"net/http"
	"strings"
)

func RelationHandler(w http.ResponseWriter, r *http.Request) {
	relationID := strings.TrimPrefix(r.URL.Path, "/relation/")
	relation := data.Relations[relationID]

	json.NewEncoder(w).Encode(relation)
}
