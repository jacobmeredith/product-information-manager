package auth

import (
	"context"

	"github.com/jacobmeredith/product-information-manager/api/internal/ports"
)

type AuthService interface {
	Login(ctx context.Context, req LoginRequest) (*LoginResponse, error)
	GeneratePassword(ctx context.Context, psswd string) (string, error)
}

type Service struct {
	ur ports.UserRepo
}

func NewService(ur ports.UserRepo) *Service {
	return &Service{
		ur: ur,
	}
}
