package handlers

import (
	"GroupieTracker/data"
	"encoding/json"
	"net/http"
	"strings"
)

func EventHandler(w http.ResponseWriter, r *http.Request) {
	eventID := strings.TrimPrefix(r.URL.Path, "/event/")
	event := data.Events[eventID]

	json.NewEncoder(w).Encode(event)
}
