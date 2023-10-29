package user

import (
	"context"
	"fmt"

	"github.com/jacobmeredith/product-information-manager/api/internal/common"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/user"
	"github.com/jacobmeredith/product-information-manager/api/pkg/errsx"
)

type CreateUserRequest struct {
	Email    string
	Password string
}

type CreateUserResponse struct {
	ID string
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error) {
	var err error
	var errs errsx.Map

	email, err := user.NewEmail(req.Email)
	if err != nil {
		errs.Set("email", err)
	}

	password, err := user.NewPassword(req.Password)
	if err != nil {
		errs.Set("password", err)
	}

	if errs != nil {
		return nil, fmt.Errorf("%w - %w", common.ErrBadRequest, errs)
	}

	user := user.NewUser(email, password)

	err = s.ur.Add(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("%w - %w", common.ErrInternal, err)
	}

	return &CreateUserResponse{ID: user.ID.String()}, nil
}
