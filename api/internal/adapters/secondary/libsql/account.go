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
