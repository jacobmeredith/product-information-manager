package user

import (
	"context"
	"fmt"

	"github.com/jacobmeredith/product-information-manager/api/internal/common"
)

type GetUserResponse struct {
	ID    string
	Email string
}

func (s *Service) GetUser(ctx context.Context, req string) (*GetUserResponse, error) {
	user, err := s.ur.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w - %w", common.ErrInternal, err)
	}

	return &GetUserResponse{ID: user.ID.String(), Email: string(user.Email)}, nil
}
