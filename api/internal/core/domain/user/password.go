package user

import (
	"errors"
)

var (
	ErrEmptyPassword = errors.New("empty email")
	ErrShortPassword = errors.New("short password")
)

type Password string

func NewPassword(password string) (Password, error) {
	if password == "" {
		return "", ErrEmptyPassword
	}

	if len(password) < 6 {
		return "", ErrShortPassword
	}

	return Password(password), nil
}
