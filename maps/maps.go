package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound          = errors.New("Couldn't find the word")
	ErrWordAlreadyExists = errors.New("Word already exists")
)

func (d Dictionary) Search(key string) (value string, err error) {
	if _, ok := d[key]; !ok {
		return "", ErrNotFound
	}
	return d[key], nil
}

func (d Dictionary) Insert(word, desc string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = desc
	case nil:
		return ErrWordAlreadyExists
	}

	return nil
}
