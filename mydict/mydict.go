package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("not found")
	errWordExists = errors.New("already exists")
	errNotExists  = errors.New("not exists")
	errCantUpdate = errors.New("can't update non-existing word")
)

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

// Delete a word
func (d Dictionary) Delete(word string) error {
	if _, err := d.Search(word); err != nil {
		return errNotExists
	}
	delete(d, word)
	return nil
}

// Update a word
func (d Dictionary) Update(word, def string) error {
	if _, err := d.Search(word); err != nil {
		return errCantUpdate
	}
	d[word] = def
	return nil
}
