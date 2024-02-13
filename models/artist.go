package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DatesResponse struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type LocationsResponse struct {
	Locations []string `json:"locations"`
}

type Artist struct {
	Id               int      `json:"id"`
	Image            string   `json:"image"`
	Name             string   `json:"name"`
	Members          []string `json:"members"`
	CreationDate     int      `json:"creationDate"`
	FirstAlbum       string   `json:"firstAlbum"`
	Locations        string   `json:"locations"`
	ConcertDates     string   `json:"concertDates"`
	Relations        string   `json:"relations"`
	LocationList     []string
	ConcertDatesList []string
}

func (a *Artist) GetLocations() error {
	resp, err := http.Get(a.Locations)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var locations LocationsResponse
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return err
	}
	a.LocationList = locations.Locations
	return nil
}

func (a *Artist) GetConcertDates() error {
	resp, err := http.Get(a.ConcertDates)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var dates DatesResponse
	err = json.Unmarshal(body, &dates)
	if err != nil {
		return err
	}

	if a.Id == dates.ID {
		a.ConcertDatesList = dates.Dates
	}

	return nil
}
