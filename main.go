package main

import "github.com/GAumala/mpd-http-proxy/mpd"

func main() {
	conn := mpd.ConnectToMPD()
	mpd.PlayNextPlaylistSong(conn)

}
