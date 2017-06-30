package mpd

import "strings"

const artistNameLinePrefix = "Artist: "
const filePathLinePrefix = "file: "
const timeLinePrefix = "Time: "
const albumNameLinePrefix = "Album: "
const songTitleLinePrefix = "Title: "
const genreLinePrefix = "Genre: "
const trackNumberLinePrefix = "Track: "

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

// ParseFindArtistResponse takes the string response from "find artists" and
// returns a LinkedListSong struct
func ParseFindArtistResponse(response LinkedListString) LinkedListSong {
	foundSongs := NewLinkedListSong()
	currentSong := new(Song)
	response.ForEach(func(s string) {
		if strings.HasPrefix(s, filePathLinePrefix) {
			parseFilepathLine(foundSongs, currentSong, s)
		} else if strings.HasPrefix(s, artistNameLinePrefix) {
			parseArtistNameLine(currentSong, s)
		} else if strings.HasPrefix(s, songTitleLinePrefix) {
			parseSongTitleLine(currentSong, s)
		} else if strings.HasPrefix(s, albumNameLinePrefix) {
			parseAlbumNameLine(currentSong, s)
		} else if strings.HasPrefix(s, timeLinePrefix) {
			parseTimeLine(currentSong, s)
		} else if strings.HasPrefix(s, genreLinePrefix) {
			parseGenreLine(currentSong, s)
		} else if strings.HasPrefix(s, trackNumberLinePrefix) {
			parseTrackNumberLine(currentSong, s)
		}
	})
	appendSongIfValid(foundSongs, currentSong)
	return foundSongs
}
