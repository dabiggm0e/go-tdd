package maps

type Dictionary map[string]string
type DictionaryError string

const (
	ErrNotFound          = DictionaryError("Couldn't find the word")
	ErrWordAlreadyExists = DictionaryError("Word already exists")
	ErrWordDoesNotExists = DictionaryError("Word does not exist in the dictionary")
	ErrWordWasntDeleted  = DictionaryError("Word was expected to be deleted")
)

func (e DictionaryError) Error() string {
	return string(e)
}

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

func (d Dictionary) Update(word, desc string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = desc
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {

	case ErrNotFound:
		{

			return ErrWordDoesNotExists
		}
	case nil:
		{
			delete(d, word)
			return nil
		}
	default:
		{
			return err
		}
	}
}
