package account

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jacobmeredith/product-information-manager/api/internal/common"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/account"
	"github.com/jacobmeredith/product-information-manager/api/pkg/errsx"
)

type InviteUserToAccountRequest struct {
	AccountId string
	Email     string
}

func (s *Service) InviteUserToAccount(ctx context.Context, req InviteUserToAccountRequest) error {
	var err error
	var errs errsx.Map

	accountId, err := uuid.Parse(req.AccountId)
	if err != nil {
		errs.Set("accountId", err)
	}

	email, err := account.NewInviteEmail(req.Email)
	if err != nil {
		errs.Set("email", err)
	}

	if errs != nil {
		return fmt.Errorf("%w - %w", common.ErrBadRequest, errs)
	}

	inv := account.NewInvite(accountId, email, account.InviteCreatedAt(time.Now()), account.InviteUpdatedAt(time.Now()), account.InviteDeletedAt(time.Now()))

	err = s.ar.InviteUserToAccount(ctx, inv.ID.String(), req.AccountId, req.Email)
	if err != nil {
		return fmt.Errorf("%w - %w", common.ErrInternal, err)
	}

	return nil
}
