package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"GroupieTracker/data"
)

func DateHandler(w http.ResponseWriter, r *http.Request) {
	dateID := strings.TrimPrefix(r.URL.Path, "/date/")
	date := data.Dates[dateID]

	json.NewEncoder(w).Encode(date)
}
