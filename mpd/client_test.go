package mpd

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

var lastWrittenRequest = ""

type fileClient struct {
	filename string
}

func (client fileClient) ReadMPDResponse() LinkedListString {
	path := path.Join(os.Getenv("GOPATH"),
		"src/github.com/GAumala/mpd-http-proxy/mpd/testData/", client.filename)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	contents := readResponseFromIOReader(file)
	file.Close()
	return contents
}

func (client fileClient) WriteMPDRequest(request string) {
	lastWrittenRequest = request
}

func TestFindArtist(t *testing.T) {
	client := fileClient{filename: "findArtistResponse1.txt"}
	songs := FindArtist(client, "Bruno Mars")

	assert.Equal(t, "find artist \"Bruno Mars\"", lastWrittenRequest,
		"should write the correct find artist request")

	expectedSongs := NewLinkedListSong()
	expectedSongs.PushBack(Song{
		filepath:    "Bruno Mars/24K Magic/01 24K Magic.m4a",
		time:        "227",
		artist:      "Bruno Mars",
		album:       "24K Magic",
		title:       "24K Magic",
		trackNumber: "1",
		genre:       "Pop",
	})
	expectedSongs.PushBack(Song{
		filepath:    "Bruno Mars/24K Magic/02 Chunky.m4a",
		time:        "187",
		artist:      "Bruno Mars",
		album:       "24K Magic",
		title:       "Chunky",
		trackNumber: "2",
		genre:       "Pop",
	})

	assert.Equal(t, expectedSongs.list, songs.list, "should parse all songs in the response text file")
}

func TestGetCurrentPlaylistResponse(t *testing.T) {
	client := fileClient{filename: "currentPlaylistResponse1.txt"}
	songs := GetCurrentPlaylistInfo(client)

	assert.Equal(t, "playlistinfo", lastWrittenRequest,
		"should write the correct playlistinfo request")

	expectedSongs := NewLinkedListSong()
	expectedSongs.PushBack(Song{
		filepath:    "Los Auténticos Decadentes/Los Reyes de la Cancion/01 Como Me Voy a Olvídar.m4a",
		time:        "205",
		artist:      "Los Auténticos Decadentes",
		album:       "Los Reyes de la Cancion",
		title:       "Como Me Voy a Olvídar",
		trackNumber: "1",
		genre:       "Música latina",
		id:          "1",
	})
	expectedSongs.PushBack(Song{
		filepath:    "Los Auténticos Decadentes/Los Reyes de la Cancion/02 No Puedo.m4a",
		time:        "188",
		artist:      "Los Auténticos Decadentes",
		album:       "Los Reyes de la Cancion",
		title:       "No Puedo",
		trackNumber: "2",
		genre:       "Música latina",
		id:          "2",
	})

	assert.Equal(t, expectedSongs.list, songs.list,
		"should parse all songs in the response text file")
}

func TestAddSongToCurrentPlaylistResponse(t *testing.T) {
	client := fileClient{filename: "addSongToCurrentPlaylistResponse1.txt"}
	id := AddSongToCurrentPlaylist(client, "path/to/song.ogg")

	assert.Equal(t, "addid \"path/to/song.ogg\"", lastWrittenRequest,
		"should write the correct addid request")
	assert.Equal(t, "5", id, "should parse the id number in the response text")
}
