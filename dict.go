package main

type Dictionary map[string]string

const (
	ErrNotFound    = DictionaryErr("could not find the word")
	ErrExistingKey = DictionaryErr("key exists in the map")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	if err != nil {
		return err
	}
	delete(d, key)
	return nil
}

func (d Dictionary) Search(keyword string) (string, error) {
	definition, ok := d[keyword]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrExistingKey
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, newValue string) error {
	_, err := d.Search(key)
	if err != nil {
		return err
	}
	d[key] = newValue
	return nil
}
