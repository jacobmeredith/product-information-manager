package user

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Email    Email
	Password Password
}

func NewUser(email Email, password Password) User {
	return User{
		ID:       uuid.New(),
		Email:    email,
		Password: password,
	}
}
