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

// Client is an interface for all mpd clients, implements methods to send
// requests to the mpd daemon and read responses.
type Client interface {
	ReadMPDResponse() LinkedListString
	WriteMPDRequest(request string)
}

// TCPClient is a Client implementation that manages a TCP connection to the
// mpd daemon
type TCPClient struct {
	conn net.Conn
}

func readResponseFromIOReader(ioReader io.Reader) LinkedListString {
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

// ReadMPDResponse reads a mpd response from the TCP connection opened in
// client and returns a LinkedListString with all the received lines.
func (client TCPClient) ReadMPDResponse() LinkedListString {
	return readResponseFromIOReader(client.conn)
}

// WriteMPDRequest writes a string with a mpd request to the TCP connection
// opened in client .
func (client TCPClient) WriteMPDRequest(request string) {
	fmt.Fprintf(client.conn, request+"\n")
}

// ConnectToMPD connects to the mpd service running in localhost and returns
// the connection object
func (client *TCPClient) ConnectToMPD() {
	conn, err := net.Dial("tcp", "127.0.0.1:6600")
	if err != nil {
		log.Fatal(err)
	}
	(*client).conn = conn
	response := client.ReadMPDResponse()
	responseString := response.String()
	if !strings.HasPrefix(responseString, "OK MPD") {
		mpdError := errors.New("Could not connect to MPD. got: \n" + responseString)
		log.Fatal(mpdError)
	}
}

// GetAllSongs calls mpd command "list" with "title" as type
func GetAllSongs(client Client) {
	client.WriteMPDRequest("list title")
	response := client.ReadMPDResponse()
	fmt.Println(response)
}

// GetAllAlbums calls mpd command "list" with "album" as type
func GetAllAlbums(client Client) {
	client.WriteMPDRequest("list album")
	response := client.ReadMPDResponse()
	fmt.Println(response)
}

// GetAllArtists calls mpd command "list" with "artist" as type
func GetAllArtists(client Client) {
	client.WriteMPDRequest("list artist")
	response := client.ReadMPDResponse()
	fmt.Println(response)
}

// FindArtist calls mpd command "find" with "artist" as type and returns a
// LinkedListSong struct with all the songs that match the query
func FindArtist(client Client, artistquery string) LinkedListSong {
	client.WriteMPDRequest("find artist \"" + artistquery + "\"")
	response := client.ReadMPDResponse()
	return ParseSongListResponse(response)
}

// GetCurrentPlaylistInfo calls mpd command "playlistinfo" and returns a
// LinkedListSong struct with all the songs in the current playlist
func GetCurrentPlaylistInfo(client Client) LinkedListSong {
	client.WriteMPDRequest("playlistinfo")
	response := client.ReadMPDResponse()
	return ParseSongListResponse(response)
}

// AddSongToCurrentPlaylist calls mpd command "addid" and returns the id
// of the new song added to the playlist
func AddSongToCurrentPlaylist(client Client, songURI string) string {
	client.WriteMPDRequest("addid \"" + songURI + "\"")
	response := client.ReadMPDResponse()
	return ParseIDResponse(response)
}

// UpdateCollection calls mpd command "update" with no additional arguments
func UpdateCollection(client Client) {
	client.WriteMPDRequest("update")
	client.ReadMPDResponse()
}

// StartPlaylist calls mpd command "play"
func StartPlaylist(client Client) {
	client.WriteMPDRequest("play")
	client.ReadMPDResponse()
}

// StopPlaylist calls mpd command "stop"
func StopPlaylist(client Client) {
	client.WriteMPDRequest("stop")
	client.ReadMPDResponse()
}

// TogglePlayback calls mpd command "pause 1" if pause is true, otherwise "pause 0"
func TogglePlayback(client Client, pause bool) {
	if pause {
		client.WriteMPDRequest("pause 1")
	} else {
		client.WriteMPDRequest("pause 0")
	}
	client.ReadMPDResponse()
}

// PlayNextPlaylistSong calls mpd command "next"
func PlayNextPlaylistSong(client Client) {
	client.WriteMPDRequest("next")
	client.ReadMPDResponse()
}

// PlayPreviousPlaylistSong calls mpd command "next"
func PlayPreviousPlaylistSong(client Client) {
	client.WriteMPDRequest("previous")
	client.ReadMPDResponse()
}
