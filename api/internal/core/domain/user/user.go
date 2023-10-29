package user

import "github.com/google/uuid"

type Email string

type User struct {
	ID    uuid.UUID
	Email Email
}
