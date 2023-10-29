package libsql

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/user"
)

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrUnknown       = errors.New("unknown error")
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) Add(ctx context.Context, u user.User) error {
	_, err := ur.db.ExecContext(ctx, "INSERT INTO user (id, email, password) VALUES (?, ?, ?)", u.ID, u.Email, u.Password)

	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "UNIQUE constraint failed") {
		return ErrAlreadyExists
	}

	return ErrUnknown
}
