package user

import (
	"context"

	"github.com/jacobmeredith/product-information-manager/api/internal/ports"
)

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error)
}

type Service struct {
	ur ports.UserRepo
}

func NewService(ur ports.UserRepo) *Service {
	return &Service{
		ur: ur,
	}
}
