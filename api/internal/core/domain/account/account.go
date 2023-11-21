package account

import "github.com/google/uuid"

type Account struct {
	ID   uuid.UUID
	Name Name
}

func NewAccount(name Name) Account {
	return Account{
		ID:   uuid.New(),
		Name: name,
	}
}
