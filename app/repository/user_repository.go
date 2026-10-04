package repository

import (
	"context"
	"fmt"

	"projek-backend/app/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (
			username,
			password,
			role
		)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Role,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("gagal membuat user: %w", err)
	}

	return nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	query := `
		SELECT
			id,
			username,
			password,
			role,
			created_at
		FROM users
		WHERE username = $1
	`

	var user model.User

	err := r.db.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user tidak ditemukan")
		}

		return nil, fmt.Errorf("gagal mencari user: %w", err)
	}

	return &user, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*model.User, error) {
	query := `
		SELECT
			id,
			username,
			password,
			role,
			created_at
		FROM users
		WHERE id = $1
	`

	var user model.User

	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user dengan id %d tidak ditemukan", id)
		}

		return nil, fmt.Errorf("gagal mencari user: %w", err)
	}

	return &user, nil
}