package main

import (
	"fmt"

	"github.com/GAumala/mpd-http-proxy/mpd"
)

func main() {
	conn := mpd.ConnectToMPD()
	songs := mpd.FindArtist(conn, "tricot")
	fmt.Printf(songs.String())

}
