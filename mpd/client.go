package mpd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func readMPDResponse(ioReader io.Reader) LinkedListString {
	reader := bufio.NewReader(ioReader)
	response := NewLinkedListString()
	for {
		newline, err := reader.ReadString('\n')
		if err != nil {
			// handle error
			log.Fatal(err)
		}
		response.PushBack(newline)
		isLastLine := strings.HasPrefix(newline, "OK") || strings.HasPrefix(newline, "ACK")
		if isLastLine {
			return response
		}
	}
}

// ConnectToMPD connects to the mpd service running in localhost and returns
// the connection object
func ConnectToMPD() *net.Conn {
	conn, err := net.Dial("tcp", "127.0.0.1:6600")
	if err != nil {
		log.Fatal(err)
	}

	response := readMPDResponse(conn)
	responseString := response.String()
	if !strings.HasPrefix(responseString, "OK MPD") {
		mpdError := errors.New("Could not connect to MPD. got: \n" + responseString)
		log.Fatal(mpdError)
	}

	return &conn
}

func writeMPDRequest(conn *net.Conn, request string) {
	fmt.Fprintf(*conn, request+"\n")
}

// GetAllSongs Calls mpd command "list" with "title" as type
func GetAllSongs(conn *net.Conn) {
	writeMPDRequest(conn, "list title")
	response := readMPDResponse(*conn)
	fmt.Println(response)
}

// GetAllAlbums Calls mpd command "list" with "album" as type
func GetAllAlbums(conn *net.Conn) {
	writeMPDRequest(conn, "list album")
	response := readMPDResponse(*conn)
	fmt.Println(response)
}

// GetAllArtists Calls mpd command "list" with "artist" as type
func GetAllArtists(conn *net.Conn) {
	writeMPDRequest(conn, "list artist")
	response := readMPDResponse(*conn)
	fmt.Println(response)
}

// FindArtist Calls mpd command "find" with "artist" as type
func FindArtist(conn *net.Conn, artistquery string) LinkedListSong {
	writeMPDRequest(conn, "find artist "+artistquery)
	response := readMPDResponse(*conn)
	return ParseSongListResponse(response)
}

// GetCurrentPlaylistInfo Calls mpd command "playlistinfo" and returns a
// LinkedListSong struct with all the songs in the current playlist
func GetCurrentPlaylistInfo(conn *net.Conn) LinkedListSong {
	writeMPDRequest(conn, "playlistinfo")
	response := readMPDResponse(*conn)
	return ParseSongListResponse(response)
}

// AddSongToCurrentPlaylist Calls mpd command "addid" and returns the id
// of the new song added to the playlist
func AddSongToCurrentPlaylist(conn *net.Conn, songURI string) string {
	writeMPDRequest(conn, "addid \""+songURI+"\"")
	response := readMPDResponse(*conn)
	return ParseIDResponse(response)
}

// UpdateCollection Calls mpd command "update" with no additional arguments
func UpdateCollection(conn *net.Conn) {
	writeMPDRequest(conn, "update")
	readMPDResponse(*conn)
}

// StartPlaylist Calls mpd command "play"
func StartPlaylist(conn *net.Conn) {
	writeMPDRequest(conn, "play")
	readMPDResponse(*conn)
}

// StopPlaylist Calls mpd command "stop"
func StopPlaylist(conn *net.Conn) {
	writeMPDRequest(conn, "stop")
	readMPDResponse(*conn)
}

// PlayNextPlaylistSong Calls mpd command "next"
func PlayNextPlaylistSong(conn *net.Conn) {
	writeMPDRequest(conn, "next")
	readMPDResponse(*conn)
}

// PlayPreviousPlaylistSong Calls mpd command "next"
func PlayPreviousPlaylistSong(conn *net.Conn) {
	writeMPDRequest(conn, "previous")
	readMPDResponse(*conn)
}
