package mpd

import "strings"

const artistNameLinePrefix = "Artist: "
const filePathLinePrefix = "file: "
const timeLinePrefix = "Time: "
const albumNameLinePrefix = "Album: "
const songTitleLinePrefix = "Title: "
const genreLinePrefix = "Genre: "
const trackNumberLinePrefix = "Track: "
const idLinePrefix = "Id: "

func appendSongIfValid(foundSongs LinkedListSong, song *Song) {
	if len(song.filepath) > 0 {
		foundSongs.PushBack(*song)
		*song = *new(Song)
	}
}
func parseFilepathLine(foundSongs LinkedListSong, currentSong *Song, line string) {
	appendSongIfValid(foundSongs, currentSong)
	(*currentSong).filepath = line[len(filePathLinePrefix) : len(line)-1]
}

func parseIDLine(currentSong *Song, line string) {
	(*currentSong).id = line[len(idLinePrefix) : len(line)-1]
}

func parseArtistNameLine(currentSong *Song, line string) {
	(*currentSong).artist = line[len(artistNameLinePrefix) : len(line)-1]
}

func parseTimeLine(currentSong *Song, line string) {
	(*currentSong).time = line[len(timeLinePrefix) : len(line)-1]
}

func parseAlbumNameLine(currentSong *Song, line string) {
	(*currentSong).album = line[len(albumNameLinePrefix) : len(line)-1]
}

func parseSongTitleLine(currentSong *Song, line string) {
	(*currentSong).title = line[len(songTitleLinePrefix) : len(line)-1]
}

func parseGenreLine(currentSong *Song, line string) {
	(*currentSong).genre = line[len(genreLinePrefix) : len(line)-1]
}

func parseTrackNumberLine(currentSong *Song, line string) {
	(*currentSong).trackNumber = line[len(trackNumberLinePrefix) : len(line)-1]
}

// ParseSongListResponse takes the string response from "find artists" and
// returns a LinkedListSong struct
func ParseSongListResponse(response LinkedListString) LinkedListSong {
	foundSongs := NewLinkedListSong()
	currentSong := new(Song)
	response.ForEach(func(s string) {
		switch {
		case strings.HasPrefix(s, filePathLinePrefix):
			parseFilepathLine(foundSongs, currentSong, s)
			break
		case strings.HasPrefix(s, idLinePrefix):
			parseIDLine(currentSong, s)
			break
		case strings.HasPrefix(s, artistNameLinePrefix):
			parseArtistNameLine(currentSong, s)
			break
		case strings.HasPrefix(s, songTitleLinePrefix):
			parseSongTitleLine(currentSong, s)
			break
		case strings.HasPrefix(s, albumNameLinePrefix):
			parseAlbumNameLine(currentSong, s)
			break
		case strings.HasPrefix(s, timeLinePrefix):
			parseTimeLine(currentSong, s)
			break
		case strings.HasPrefix(s, genreLinePrefix):
			parseGenreLine(currentSong, s)
			break
		case strings.HasPrefix(s, trackNumberLinePrefix):
			parseTrackNumberLine(currentSong, s)
			break
		}
	})
	appendSongIfValid(foundSongs, currentSong)
	return foundSongs
}

// ParseIDResponse takes the string response from "addid" command and
// returns a string with the id of the added song
func ParseIDResponse(response LinkedListString) string {
	firstLine := response.Front()
	song := new(Song)
	parseIDLine(song, firstLine)
	return song.id
}
