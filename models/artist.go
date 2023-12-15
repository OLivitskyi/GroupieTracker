package models

type Artist struct {
	ArtistID   string
	Name       string
	Image      string
	StartYear  int
	FirstAlbum string
	Members    []string
}
