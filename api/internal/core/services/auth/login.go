package auth

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jacobmeredith/product-information-manager/api/internal/common"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/user"
	"github.com/jacobmeredith/product-information-manager/api/pkg/errsx"
)

type CustomClaims struct {
	userId string
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Email    string
	Password string
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (string, error) {
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
		return "", fmt.Errorf("%w - %w", common.ErrBadRequest, errs)
	}

	user, err := s.VerifyPassword(ctx, VerifyPasswordRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return "", common.ErrInternal
	}

	if user == nil {
		return "", errors.New("invalid credentials")
	}

	userId := user.ID.String()

	claims := CustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", common.ErrInternal
	}

	return signedToken, nil
}
