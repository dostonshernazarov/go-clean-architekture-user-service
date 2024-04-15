// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github/go-clean-template/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Post -. 
	User interface {
		CreateUser(context.Context, *entity.User) (*entity.User, error)
		GetUser(context.Context, string) (*entity.User, error) 
		UpdateUser(context.Context, *entity.User) (*entity.User, error)
		DeleteUser(context.Context, string) (error) 
	}

	// UserRepo -. 
	UserRepo interface {
		Create(context.Context, *entity.User) (*entity.User, error)
		Get(context.Context, string) (*entity.User, error) 
		Update(context.Context, *entity.User) (*entity.User, error)
		Delete(context.Context, string) (error) 
	}
)
