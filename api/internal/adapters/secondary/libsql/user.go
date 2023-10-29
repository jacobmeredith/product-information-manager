package libsql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/user"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) Add(ctx context.Context, u user.User) error {
	_, err := ur.db.ExecContext(ctx, "INSERT INTO user (id, email, password) VALUES (?, ?, ?)", u.ID, u.Email, u.Password)
	fmt.Println(err)
	return err
}
