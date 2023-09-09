package valueobjects

import (
	"errors"
	"net/url"
)

var (
	ErrEmptyInputUrl   = errors.New("empty url")
	ErrInvalidInputUrl = func(err error) error {
		if err != nil {
			return errors.New("invalid url: " + err.Error())
		}

		return errors.New("invalid url")
	}
)

type Address struct {
	Url url.URL
}

func NewAddress(input string) (Address, error) {
	if input == "" {
		return Address{}, ErrEmptyInputUrl
	}

	u, err := url.ParseRequestURI(input)
	if err != nil {
		return Address{}, ErrInvalidInputUrl(err)
	}

	return Address{Url: *u}, nil
}
