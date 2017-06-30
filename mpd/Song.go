package mpd

import "fmt"

// Song is a model for each song returned by mpd's "find" command
type Song struct {
	id          string
	filepath    string
	title       string
	album       string
	artist      string
	time        string
	trackNumber string
	genre       string
}

func (s Song) String() string {
	return fmt.Sprintf("{ id:\"%s\", filepath:\"%s\", title:\"%s\", album:\"%s\", "+
		"artist:\"%s\", time:\"%s\", trackNumber:\"%s\", genre:\"%s\" }", s.id,
		s.filepath, s.title, s.album, s.artist, s.time, s.trackNumber, s.genre)
}
