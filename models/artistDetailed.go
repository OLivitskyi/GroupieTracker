package models

type ArtistDetailed struct {
	Artist
	Locations    []string
	ConcertDates []string
}
