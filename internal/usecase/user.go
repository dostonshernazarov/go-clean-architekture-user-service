package usecase

import (
	"context"
	"fmt"

	"github/go-clean-template/internal/entity"
)

// UserUseCase -.
type UserUseCase struct {
	repo UserRepo
}

// New -.
func New(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo:   r,
	}
}

// CreateUser 
func (uc *UserUseCase) CreateUser(ctx context.Context, req *entity.User) (*entity.User, error) {
	Users, err := uc.repo.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - Create - s.repo: %w", err)
	}

	return Users, nil
}

// Get User
func (uc *UserUseCase) GetUser(ctx context.Context, id string) (*entity.User, error) {
	User, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - Get - p.repo: %w", err)
	}

	return User, nil
}

// Update User
func (p *UserUseCase) UpdateUser(ctx context.Context, req *entity.User) (*entity.User, error) {
	User, err := p.repo.Update(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("UserUseCase - Update - p.repo: %w", err)
	}

	return User, nil
}

// Delete User
func (p *UserUseCase) DeleteUser(ctx context.Context, id string) error {
	err := p.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("UserUseCase - Delete - p.repo: %w", err)
	}

	return nil
}