package repository

import (
	"context"
	"fmt"

	"projek-backend/app/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VehicleCategoryRepository struct {
	db *pgxpool.Pool
}

func NewVehicleCategoryRepository(db *pgxpool.Pool) *VehicleCategoryRepository {
	return &VehicleCategoryRepository{
		db: db,
	}
}

func (r *VehicleCategoryRepository) Create(ctx context.Context, category *model.VehicleCategory) error {
	query := `
		INSERT INTO vehicle_categories (
			name,
			description
		)
		VALUES ($1, $2)
		RETURNING id, created_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		category.Name,
		category.Description,
	).Scan(
		&category.ID,
		&category.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("gagal membuat kategori kendaraan: %w", err)
	}

	return nil
}

func (r *VehicleCategoryRepository) FindAll(ctx context.Context) ([]model.VehicleCategory, error) {
	query := `
		SELECT
			id,
			name,
			description,
			created_at
		FROM vehicle_categories
		ORDER BY id DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil kategori kendaraan: %w", err)
	}
	defer rows.Close()

	var categories []model.VehicleCategory

	for rows.Next() {
		var category model.VehicleCategory

		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("gagal membaca data kategori kendaraan: %w", err)
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("terjadi kesalahan saat membaca kategori kendaraan: %w", err)
	}

	return categories, nil
}

func (r *VehicleCategoryRepository) FindByID(ctx context.Context, id int) (*model.VehicleCategory, error) {
	query := `
		SELECT
			id,
			name,
			description,
			created_at
		FROM vehicle_categories
		WHERE id = $1
	`

	var category model.VehicleCategory

	err := r.db.QueryRow(ctx, query, id).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("kategori kendaraan dengan id %d tidak ditemukan", id)
		}

		return nil, fmt.Errorf("gagal mengambil kategori kendaraan: %w", err)
	}

	return &category, nil
}

func (r *VehicleCategoryRepository) Update(ctx context.Context, category *model.VehicleCategory) error {
	query := `
		UPDATE vehicle_categories
		SET
			name = $1,
			description = $2
		WHERE id = $3
	`

	result, err := r.db.Exec(
		ctx,
		query,
		category.Name,
		category.Description,
		category.ID,
	)

	if err != nil {
		return fmt.Errorf("gagal mengubah kategori kendaraan: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("kategori kendaraan dengan id %d tidak ditemukan", category.ID)
	}

	return nil
}

func (r *VehicleCategoryRepository) Delete(ctx context.Context, id int) error {
	query := `
		DELETE FROM vehicle_categories
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus kategori kendaraan: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("kategori kendaraan dengan id %d tidak ditemukan", id)
	}

	return nil
}