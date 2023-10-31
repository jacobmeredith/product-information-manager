package auth

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) GeneratePassword(ctx context.Context, psswd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(psswd), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
