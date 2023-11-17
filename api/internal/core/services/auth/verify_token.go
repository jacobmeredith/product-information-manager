package auth

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jacobmeredith/product-information-manager/api/internal/core/domain/user"
)

func (s *Service) VerifyToken(ctx context.Context, token string) (*user.User, error) {
	pt, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(parsedToken *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := pt.Claims.(*CustomClaims)
	if !ok || !pt.Valid {
		return nil, errors.New("Invalid token")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, errors.New("Token expired")
	}

	return &claims.User, nil
}
