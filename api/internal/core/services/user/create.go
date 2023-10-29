package user

import (
	"context"
	"errors"

	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/user"
)

type CreateUserRequest struct {
	Email    string
	Password string
}

type CreateUserResponse struct {
	ID string
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error) {
	email, err := user.NewEmail(req.Email)
	if err != nil {
		return nil, errors.Join(err, errors.New("invalid email supplied"))
	}

	password, err := user.NewPassword(req.Password)
	if err != nil {
		return nil, errors.Join(err, errors.New("invalid password supplied"))
	}

	user := user.NewUser(email, password)

	err = s.ur.Add(ctx, user)
	if err != nil {
		return nil, errors.Join(err, errors.New("failed to save user"))
	}

	return &CreateUserResponse{ID: user.ID.String()}, nil
}
