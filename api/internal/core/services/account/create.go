package account

import (
	"context"
	"fmt"

	"github.com/jacobmeredith/product-information-manager/api/internal/common"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/account"
	"github.com/jacobmeredith/product-information-manager/api/pkg/errsx"
)

type CreateAccountRequest struct {
	Name string
}

type CreateAccountResponse struct {
	ID string
}

func (s *Service) CreateAccount(ctx context.Context, req CreateAccountRequest) (*CreateAccountResponse, error) {
	var err error
	var errs errsx.Map

	name, err := account.NewName(req.Name)
	if err != nil {
		errs.Set("email", err)
	}

	if errs != nil {
		return nil, fmt.Errorf("%w - %w", common.ErrBadRequest, errs)
	}

	u := account.NewAccount(name)
	// Repository add to database

	return &CreateAccountResponse{ID: u.ID.String()}, nil
}
