package main

const (
	ErrDuplicate = DictionaryErr("duplicate entry")
	ErrNotFound  = DictionaryErr("could not find the word sought")
	ErrNoExist   = DictionaryErr("cannot update word because it doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

/*var ErrDuplicate = errors.New("duplicate entry")
var ErrNotFound = errors.New("could not find the word sought")*/

type Dictionary map[string]string

type MapSearch interface {
	Search(dictionary Dictionary, key string) string
}

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]

}

func (d *Dictionary) Search(key string) (string, error) {
	value, ok := (*d)[key]
	if ok {
		return value, nil
	} else {
		return "", ErrNotFound
	}

}

func (d *Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		(*d)[key] = value
		return nil
	case nil:
		return ErrDuplicate
	default:
		return err

	}
}

func (d *Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrNoExist
	case nil:
		(*d)[key] = value
		return nil
	default:
		return err

	}
}
