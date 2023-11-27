package ports

import (
	"context"

	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/account"
)

type AccountRepo interface {
	Add(ctx context.Context, a account.Account) error
	AddUserToAccount(ctx context.Context, role string, userId string, accountId string) error
	Get(ctx context.Context, id string) (*account.Account, error)
}
