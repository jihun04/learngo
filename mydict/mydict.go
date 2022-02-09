package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("not found")
var errWordExists = errors.New("already exists")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to dictionary
func (d Dictionary) Add(word, def string) error {
	if _, err := d.Search(word); err == nil {
		return errWordExists
	}
	d[word] = def
	return nil
}
