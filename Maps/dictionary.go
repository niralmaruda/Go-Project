package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrorNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrorWordExists = DictionaryErr("cannot add word because it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	defination, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return defination, nil
}

func (d Dictionary) Add(word, defination string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = defination
	case nil:
		return ErrorWordExists
	default:
		return err
	}
	return nil
}
