package libsql

import (
	"context"
	"database/sql"
	"strings"

	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/account"
)

type AccountRepo struct {
	db *sql.DB
}

func NewAccountRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{db: db}
}

func (ur *AccountRepo) Add(ctx context.Context, u account.Account) error {
	_, err := ur.db.ExecContext(ctx, "INSERT INTO account (id, name) VALUES (?, ?)", u.ID, u.Name)

	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "UNIQUE constraint failed") {
		return ErrAlreadyExists
	}

	return ErrUnknown
}

func (ur *AccountRepo) InviteUserToAccount(ctx context.Context, id string, accountId string, email string) error {
	_, err := ur.db.ExecContext(ctx, "INSERT INTO account_user_invite (id, account_id, email)", id, accountId, email)
	if err != nil {
		return err
	}

	return nil
}

func (ur *AccountRepo) GetUserInvite(ctx context.Context, id string) (*account.Invite, error) {
	invite := new(account.Invite)
	row := ur.db.QueryRowContext(ctx, "SELECT id, account_id, email, created_at, updated_at, deleted_at FROM account_user_invite WHERE id=?", id)
	err := row.Scan(&invite.ID, &invite.AccountID, &invite.Email, &invite.CreatedAt, &invite.UpdatedAt, &invite.DeletedAt)
	if err == nil {
		return invite, nil
	}

	return nil, ErrUnknown
}

func (ur *AccountRepo) InvalidateUserInvite(ctx context.Context, id string) error {
	_, err := ur.db.ExecContext(ctx, "UPDATE account_user_invite SET deleted_at = CURRENT_TIMESTAMP WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (ur *AccountRepo) AddUserToAccount(ctx context.Context, role string, userId string, accountId string) error {
	_, err := ur.db.ExecContext(ctx, "INSERT INTO account_user (role, account_id, user_id) VALUES (?, ?, ?)", role, accountId, userId)

	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "UNIQUE constraint failed") {
		return ErrAlreadyExists
	}

	return ErrUnknown
}

func (ur *AccountRepo) Get(ctx context.Context, id string) (*account.Account, error) {
	account := new(account.Account)
	row := ur.db.QueryRowContext(ctx, "SELECT id, name FROM account WHERE id=?", id)
	err := row.Scan(&account.ID, &account.Name)
	if err == nil {
		return account, nil
	}

	return nil, ErrUnknown
}
