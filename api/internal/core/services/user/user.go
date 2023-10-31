package user

import (
	"context"

	"github.com/jacobmeredith/product-information-manager/api/internal/core/services/auth"
	"github.com/jacobmeredith/product-information-manager/api/internal/ports"
)

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error)
	GetUser(ctx context.Context, req string) (*GetUserResponse, error)
}

type Service struct {
	as auth.AuthService
	ur ports.UserRepo
}

func NewService(as auth.AuthService, ur ports.UserRepo) *Service {
	return &Service{
		as: as,
		ur: ur,
	}
}
