package auth

import (
	"context"

	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type VerifyPasswordRequest struct {
	Email    user.Email
	Password user.Password
}

func (s *Service) VerifyPassword(ctx context.Context, req VerifyPasswordRequest) (*user.User, error) {
	u, err := s.ur.GetByEmail(ctx, string(req.Email))
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	return u, nil
}
