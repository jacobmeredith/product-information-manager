package account

import (
	"errors"
	"strings"
)

var (
	ErrEmptyName    = errors.New("empty name")
	ErrSpecialChars = errors.New("name should not contain any special characters")
)

type Name string

func NewName(name string) (Name, error) {
	if name == "" {
		return "", ErrEmptyName
	}

	if strings.ContainsAny(name, "!@#$%^&*()_+{}|:<>?") {
		return "", ErrSpecialChars
	}

	return Name(name), nil
}
