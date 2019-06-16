package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("Couldn't find the word")

func (d Dictionary) Search(key string) (value string, err error) {
	if _, ok := d[key]; !ok {
		return "", ErrNotFound
	}
	return d[key], nil
}
