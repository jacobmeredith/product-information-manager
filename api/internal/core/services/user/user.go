package user

import "context"

type UserService interface {
	CreateUser(ctx context.Context, req CreateUserRequest) (*CreateUserResponse, error)
}

type Service struct {
}

func NewService() *Service {
	return &Service{}
}
