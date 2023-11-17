package auth

import (
	"context"

	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/user"
	"github.com/jacobmeredith/product-information-manager/api/internal/ports"
)

type AuthService interface {
	Login(ctx context.Context, req LoginRequest) (string, error)
	GeneratePassword(ctx context.Context, psswd string) (string, error)
	VerifyPassword(ctx context.Context, req VerifyPasswordRequest) (*user.User, error)
	VerifyToken(ctx context.Context, token string) (*user.User, error)
}

type Service struct {
	ur ports.UserRepo
}

func NewService(ur ports.UserRepo) *Service {
	return &Service{
		ur: ur,
	}
}
