package ports

import (
	"context"

	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/account"
)

type AccountRepo interface {
	Add(ctx context.Context, a account.Account) error
}
