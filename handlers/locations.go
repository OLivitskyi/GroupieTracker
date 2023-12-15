package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"GroupieTracker/data"
)

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	locationID := strings.TrimPrefix(r.URL.Path, "/location/")
	location := data.Locations[locationID]

	json.NewEncoder(w).Encode(location)
}
