package mpd

import (
	"bytes"
	"container/list"
)

//LinkedListString is a list only for strings
type LinkedListString struct {
	list *list.List
}

// NewLinkedListString returns an initialized LinkedListString.
func NewLinkedListString() LinkedListString {
	return LinkedListString{list: list.New()}
}

// Front returns the first string of list l or nil.
func (l LinkedListString) Front() string {
	frontElement := l.list.Front()
	return frontElement.Value.(string)
}

// PushBack inserts a new string s at the back of the list l and returns
//a reference to s.
func (l LinkedListString) PushBack(s string) string {
	l.list.PushBack(s)
	return s
}

// ForEach executes the callback function for each string in the list l.
func (l LinkedListString) ForEach(callback func(s string)) {
	for e := l.list.Front(); e != nil; e = e.Next() {
		stringValue := e.Value.(string)
		callback(stringValue)
	}
}

// String concatenates all elements of the list to a single string
func (l LinkedListString) String() string {
	var buffer bytes.Buffer
	l.ForEach(func(s string) {
		buffer.WriteString(s)
	})
	return buffer.String()
}
