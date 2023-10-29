package user

import (
	"errors"
)

var (
	ErrEmptyPassword = errors.New("empty email")
	ErrShortPassword = errors.New("password must be more than 6 characters")
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
