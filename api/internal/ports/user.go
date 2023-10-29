package ports

import "github.com/jacobmeredith/product-information-manager/api/internal/core/domain/user"

type UserRepo interface {
	Add() (user.User, error)
	Get() (user.User, error)
	Update() (user.User, error)
	Delete() (user.User, error)
	GetAll() ([]user.User, error)
}
