package main

import (
	"GroupieTracker/models"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var artists []models.Artist
		err = json.Unmarshal(body, &artists)
		if err != nil {
			log.Fatal(err)
		}

		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		err = tmpl.Execute(w, artists)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/artist/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/artist/")
		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var artist models.Artist
		err = json.Unmarshal(body, &artist)
		if err != nil {
			log.Fatal(err)
		}

		err = artist.GetLocations()
		if err != nil {
			log.Fatal(err)
		}

		err = artist.GetConcertDates()
		if err != nil {
			log.Fatal(err)
		}
		err = artist.GetConcertDates()
		if err != nil {
			log.Fatal(err)
		}

		tmpl := template.Must(template.ParseFiles("templates/artist.html"))
		err = tmpl.Execute(w, artist)
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			r.ParseForm()

			id := r.Form.Get("query")
			if id == "" {
				http.Error(w, "The search field should not be empty.", http.StatusBadRequest)
				return
			}

			if _, err := strconv.Atoi(id); err != nil {
				http.Error(w, "Invalid ID format. Please enter a number.", http.StatusBadRequest)
				return
			}

			resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
			if err != nil {
				http.Error(w, "An error occurred while processing your request.", http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)
			var artist models.Artist
			err = json.Unmarshal(body, &artist)
			if err != nil {
				http.Error(w, "Artist not found.", http.StatusNotFound)
				return
			}

			err = artist.GetLocations()
			if err != nil {
				http.Error(w, "An error occurred while processing your request.", http.StatusInternalServerError)
				return
			}

			err = artist.GetConcertDates()
			if err != nil {
				http.Error(w, "An error occurred while processing your request.", http.StatusInternalServerError)
				return
			}

			tmpl := template.Must(template.ParseFiles("templates/searchResults.html"))
			err = tmpl.Execute(w, artist)
			if err != nil {
				http.Error(w, "An error occurred while processing your request.", http.StatusInternalServerError)
				return
			}
		}
	})
	log.Fatal(http.ListenAndServe(":8090", nil))
}
