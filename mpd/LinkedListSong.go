package mpd

import (
	"bytes"
	"container/list"
)

//LinkedListSong is a list only for song structs
type LinkedListSong struct {
	list *list.List
}

// NewLinkedListSong returns an initialized LinkedListSong.
func NewLinkedListSong() LinkedListSong {
	return LinkedListSong{list: list.New()}
}

// Len returns the number of elements of LinkedListSong l. The complexity is O(1).
func (l LinkedListSong) Len() int {
	return l.list.Len()
}

// PushBack inserts a new song s at the back of the list l and returns
//a copy of s.
func (l LinkedListSong) PushBack(s Song) Song {
	l.list.PushBack(s)
	return s
}

// ForEach executes the callback function for each song in the list l.
func (l LinkedListSong) ForEach(callback func(s Song)) {
	for e := l.list.Front(); e != nil; e = e.Next() {
		songValue := e.Value.(Song)
		callback(songValue)
	}
}

// String concatenates all songs of the list to a single string
func (l LinkedListSong) String() string {
	var buffer bytes.Buffer
	l.ForEach(func(s Song) {
		buffer.WriteString(s.String())
	})
	return buffer.String()
}
