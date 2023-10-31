package auth

import (
	"context"

	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/user"
)

type LoginRequest struct {
	Email    user.Email
	Password user.Password
}

type LoginResponse struct {
	Token string
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	return nil, nil
}
