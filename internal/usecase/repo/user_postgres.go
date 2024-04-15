package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github/go-clean-template/internal/entity"
	"github/go-clean-template/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// Create user -.
func (p *UserRepo) Create(ctx context.Context, req *entity.User) (*entity.User, error) {
	if req.Id == "" {
		req.Id = uuid.New().String()
	}

	query, args, err := p.Builder.Insert("user_info").
		Columns(`
			id,
			full_name,
			username,
			email,
			password,
			bio,
			website,
			created_at
		`).
		Values(
			req.Id, req.FullName, req.Username, req.Email,
			req.Password, req.Bio, req.Website,  time.Now()).Suffix(
		`RETURNING created_at, updated_at`,
	).ToSql()
	if err != nil {
		return nil, fmt.Errorf("PostRepo - CreateUser - p.Builder: %w", err)
	}
	var (
		createdAt time.Time
		updatedAt sql.NullTime
	)


	row := p.Pool.QueryRow(ctx, query, args...)
	if err := row.Scan(&createdAt, &updatedAt); err != nil {
		return nil, fmt.Errorf("PostRepo - CreatePost row.Scan: %v", err)
	}


	return req, nil
}

// Get User -.
func (p *UserRepo) Get(ctx context.Context, id string) (*entity.User, error) {
	query := p.Builder.
		Select(`
		id,
		full_name,
		username,
		email,
		password,
		bio,
		website,
		created_at,
		updated_at
		`).
		From("user_info")

	if id != "" {
		query = query.Where(squirrel.Eq{"id": id})
	} else {
		return nil, fmt.Errorf("id is required")
	}

	q, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetUser - p.Builder: %w", err)
	}

	var (
		user      entity.User
		createdAt time.Time
		updatedAt sql.NullTime
	)

	row := p.Pool.QueryRow(ctx, q, args...)
	if err := row.Scan(&user.Id, &user.FullName, &user.Username,
		&user.Email, &user.Password, &user.Bio, &user.Website, &createdAt, &updatedAt); err != nil {
		return nil, fmt.Errorf("UserRepo - GetUser row.Scan: %w", err)
	}


	return &user, nil

}


// Update User -.
func (p *UserRepo) Update(ctx context.Context, req *entity.User) (*entity.User, error) {
	var (
		updateMap = make(map[string]interface{})
		where     = squirrel.Eq{"id": req.Id}
	)

	updateMap["full_name"] = req.FullName
	updateMap["username"] = req.Username
	updateMap["email"] = req.Email
	updateMap["password"] = req.Password
	updateMap["bio"] = req.Bio
	updateMap["website"] = req.Website
	updateMap["updated_at"] = time.Now()

	query := p.Builder.Update("user_info").SetMap(updateMap).Where(where).Suffix("RETURNING created_at, updated_at")
	var (
		createdAt time.Time
		updatedAt sql.NullTime
	)

	q, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserRepo - UpdateUser - p.Builder: %w", err)
	}
	row := p.Pool.QueryRow(ctx, q, args...)
	if err := row.Scan(&createdAt, &updatedAt); err != nil {
		return nil, fmt.Errorf("UserRepo - GetUser row.Scan: %w", err)
	}


	return req, nil
}

// Delete User -,
func (p *UserRepo) Delete(ctx context.Context, id string) error {
	query := p.Builder.Delete("user_info")
	if id != "" {
		query = query.Where(squirrel.Eq{"id": id})
	} else {
		return fmt.Errorf("id is required")
	}

	q, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("UserRepo - DeleteUser - p.Builder: %w", err)
	}

	_, err = p.Pool.Exec(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("UserRepo - DeleteUser row Exec: %w", err)
	}

	return nil
}
