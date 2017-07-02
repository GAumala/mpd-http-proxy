package main

import "github.com/GAumala/mpd-http-proxy/mpd"

func main() {
	client := mpd.TCPClient{}
	client.ConnectToMPD()
	mpd.PlayNextPlaylistSong(client)

}
