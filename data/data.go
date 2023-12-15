package data

import (
	"GroupieTracker/models"
)

var Artists = map[string]*models.Artist{
	"1": &models.Artist{
		ArtistID:   "1",
		Name:       "August Burns Red",
		Image:      "august_burns_red.jpg",
		StartYear:  2003,
		FirstAlbum: "Thrill Seeker",
		Members:    []string{"Jake Luhrs", "John Benjamin Brubaker", "Brent Rambler", "Matt Greiner", "Dustin Davidson"},
	},
	"2": &models.Artist{
		ArtistID:   "2",
		Name:       "Disturbed",
		Image:      "disturbed.jpg",
		StartYear:  1996,
		FirstAlbum: "The Sickness",
		Members:    []string{"David Draiman", "Dan Donegan", "Mike Wengren", "John Moyer"},
	},
}

var Locations = map[string]*models.Location{
	"1": &models.Location{
		LocationID: "1",
		Country:    "USA",
		City:       "New York",
	},
	"2": &models.Location{
		LocationID: "2",
		Country:    "UK",
		City:       "London",
	},
}

var Dates = map[string]*models.Date{
	"1": &models.Date{
		DateID:       "1",
		ConcertDates: []string{"01-05-2022", "06-10-2023"},
	},
}

var Relations = map[string]*models.Relation{
	"1": &models.Relation{
		RelationID: "1",
		ArtistID:   "1",
		LocationID: "1",
		DateID:     "1",
	},
	"2": &models.Relation{
		RelationID: "2",
		ArtistID:   "2",
		LocationID: "2",
		DateID:     "2",
	},
}

var Events = map[string]*models.Event{
	"1": &models.Event{
		EventID:   "1",
		EventType: "concert",
		Location:  "New York",
		Timestamp: "01-05-2022 19:00",
	},
	"2": &models.Event{
		EventID:   "2",
		EventType: "concert",
		Location:  "London",
		Timestamp: "06-10-2023 20:00",
	},
}
