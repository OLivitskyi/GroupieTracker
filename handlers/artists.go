package handlers

import (
	"GroupieTracker/data"
	"encoding/json"
	"net/http"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artistID := strings.TrimPrefix(r.URL.Path, "/artist/")
	artist := data.Artists[artistID]

	json.NewEncoder(w).Encode(artist)
}
