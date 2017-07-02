package main

import (
	"flag"
	"fmt"

	"github.com/GAumala/mpd-http-proxy/mpd"
)

type cmd int

const (
	start cmd = iota
	stop
	next
	previous
	pause
	resume
	undef
)

func parseCommand() cmd {
	startFlag := flag.Bool("start", false, "Starts the current playlist")
	stopFlag := flag.Bool("stop", false, "Stops the current playlist")
	nextFlag := flag.Bool("next", false, "Plays the next song in the current playlist")
	prevFlag := flag.Bool("prev", false, "Plays the previous song in the current playlist")
	pauseFlag := flag.Bool("pause", false, "Pauses the current playlist")
	resumeFlag := flag.Bool("resume", false, "Resumes the current playlist")

	flag.Parse()

	switch {
	case *startFlag:
		return start
	case *stopFlag:
		return stop
	case *nextFlag:
		return next
	case *prevFlag:
		return previous
	case *pauseFlag:
		return pause
	case *resumeFlag:
		return resume
	}

	return undef
}

func main() {
	client := mpd.TCPClient{}
	client.ConnectToMPD()
	switch parseCommand() {
	case start:
		mpd.StartPlaylist(client)
		return
	case stop:
		mpd.StopPlaylist(client)
		return
	case next:
		mpd.PlayNextPlaylistSong(client)
		return
	case previous:
		mpd.PlayPreviousPlaylistSong(client)
		return
	case pause:
		mpd.TogglePlayback(client, true)
		return
	case resume:
		mpd.TogglePlayback(client, false)
		return
	case undef:
		fmt.Println("Unknown command. Available commands are:")
		flag.PrintDefaults()
		return
	}
}
