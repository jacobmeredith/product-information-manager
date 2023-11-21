package account

import "context"

type AccountService interface {
	CreateAccount(ctx context.Context, req CreateAccountRequest) (*CreateAccountResponse, error)
}

type Service struct {
}

func NewService() AccountService {
	return &Service{}
}
