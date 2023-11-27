package account

import (
	"context"
	"fmt"

	"github.com/jacobmeredith/product-information-manager/api/internal/common"
)

type GetAccountResponse struct {
	ID   string
	Name string
}

func (s *Service) GetAccount(ctx context.Context, req string) (*GetAccountResponse, error) {
	account, err := s.ar.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w - %w", common.ErrInternal, err)
	}

	return &GetAccountResponse{ID: account.ID.String(), Name: string(account.Name)}, nil
}
