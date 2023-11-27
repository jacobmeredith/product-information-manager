package libsql

import (
	"context"
	"database/sql"
	"fmt"
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

func (ur *AccountRepo) AddUserToAccount(ctx context.Context, role string, userId string, accountId string) error {
	_, err := ur.db.ExecContext(ctx, "INSERT INTO account_user (role, account_id, user_id) VALUES (?, ?, ?)", role, accountId, userId)

	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "UNIQUE constraint failed") {
		return ErrAlreadyExists
	}

	fmt.Println(err)

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
