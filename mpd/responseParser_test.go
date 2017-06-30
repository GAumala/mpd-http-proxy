package mpd

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readTestFile(filename string) LinkedListString {
	path := path.Join(os.Getenv("GOPATH"), "src/github.com/GAumala/mpd-http-proxy/mpd/testData/", filename)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	contents := readMPDResponse(file)
	file.Close()
	return contents
}

func TestParseFindArtistResponse(t *testing.T) {
	response := readTestFile("findArtistResponse1.txt")
	songs := ParseSongListResponse(response)

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
	response := readTestFile("currentPlaylistResponse1.txt")
	songs := ParseSongListResponse(response)

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

	assert.Equal(t, expectedSongs.list, songs.list, "should parse all songs in the response text file")
}

func TestAddSongToCurrentPlaylistResponse(t *testing.T) {
	response := readTestFile("addSongToCurrentPlaylistResponse1.txt")
	id := ParseIDResponse(response)
	assert.Equal(t, "5", id, "should parse the id number in the response text")
}
