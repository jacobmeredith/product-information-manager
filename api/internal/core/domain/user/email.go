package user

import (
	"errors"
	"strings"
)

var (
	ErrEmptyEmail   = errors.New("empty email")
	ErrInvalidEmail = errors.New("invalid email")
)

type Email string

func NewEmail(email string) (Email, error) {
	if email == "" {
		return "", ErrEmptyEmail
	}

	if strings.Contains(email, "@") == false {
		return "", ErrInvalidEmail
	}

	return Email(email), nil
}
