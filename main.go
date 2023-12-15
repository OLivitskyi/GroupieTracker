package main

import (
	"log"
	"net/http"

	"GroupieTracker/handlers"
)

func main() {
	http.HandleFunc("/artist/", handlers.ArtistHandler)
	http.HandleFunc("/location/", handlers.LocationHandler)
	http.HandleFunc("/date/", handlers.DateHandler)
	http.HandleFunc("/relation/", handlers.RelationHandler)
	http.HandleFunc("/event/", handlers.EventHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
