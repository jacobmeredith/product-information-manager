package ports

import (
	"context"

	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/user"
)

type UserRepo interface {
	Add(ctx context.Context, u user.User) error
	Get(ctx context.Context, id string) (*user.User, error)
	// Update() error
	// Delete() error
	// GetAll() error
}
