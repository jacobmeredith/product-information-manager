package account

import (
	"errors"
	"strings"
)

var (
	ErrEmptyName    = errors.New("empty name")
	ErrNameLength   = errors.New("name should be 4 or more characters")
	ErrSpecialChars = errors.New("name should not contain any special characters")
)

type Name string

func NewName(name string) (Name, error) {
	if name == "" {
		return "", ErrEmptyName
	}

	if len(name) < 4 {
		return "", ErrNameLength
	}

	if strings.ContainsAny(name, "!@#$%^&*()_+{}|:<>?") {
		return "", ErrSpecialChars
	}

	return Name(name), nil
}
