package account

import (
	"context"

	"github.com/jacobmeredith/product-information-manager/api/internal/ports"
)

type AccountService interface {
	CreateAccount(ctx context.Context, req CreateAccountRequest) (*CreateAccountResponse, error)
	InviteUserToAccount(ctx context.Context, req InviteUserToAccountRequest) error
	InviteUserAccept(ctx context.Context, req InviteUserAcceptRequest) error
	GetAccount(ctx context.Context, id string) (*GetAccountResponse, error)
}

type Service struct {
	ar ports.AccountRepo
}

func NewService(ar ports.AccountRepo) AccountService {
	return &Service{
		ar: ar,
	}
}
