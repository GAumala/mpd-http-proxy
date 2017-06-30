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
	songs := ParseFindArtistResponse(response)

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
