package repository

import (
	"context"
	"fmt"

	"projek-backend/app/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type VehicleRepository struct {
	db *pgxpool.Pool
}

func NewVehicleRepository(db *pgxpool.Pool) *VehicleRepository {
	return &VehicleRepository{
		db: db,
	}
}

func (r *VehicleRepository) Create(ctx context.Context, vehicle *model.Vehicle) error {
	query := `
		INSERT INTO vehicles (
			category_id,
			name,
			brand,
			model,
			license_plate,
			year,
			price_per_day,
			status
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		vehicle.CategoryID,
		vehicle.Name,
		vehicle.Brand,
		vehicle.Model,
		vehicle.LicensePlate,
		vehicle.Year,
		vehicle.PricePerDay,
		vehicle.Status,
	).Scan(
		&vehicle.ID,
		&vehicle.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("gagal membuat kendaraan: %w", err)
	}

	return nil
}

func (r *VehicleRepository) FindAll(ctx context.Context) ([]model.Vehicle, error) {
	query := `
		SELECT
			id,
			category_id,
			name,
			brand,
			model,
			license_plate,
			year,
			price_per_day,
			status,
			created_at
		FROM vehicles
		ORDER BY id DESC
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("gagal mengambil kendaraan: %w", err)
	}
	defer rows.Close()

	var vehicles []model.Vehicle

	for rows.Next() {
		var vehicle model.Vehicle

		err := rows.Scan(
			&vehicle.ID,
			&vehicle.CategoryID,
			&vehicle.Name,
			&vehicle.Brand,
			&vehicle.Model,
			&vehicle.LicensePlate,
			&vehicle.Year,
			&vehicle.PricePerDay,
			&vehicle.Status,
			&vehicle.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("gagal membaca data kendaraan: %w", err)
		}

		vehicles = append(vehicles, vehicle)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("terjadi kesalahan saat membaca kendaraan: %w", err)
	}

	return vehicles, nil
}

func (r *VehicleRepository) FindByID(ctx context.Context, id int) (*model.Vehicle, error) {
	query := `
		SELECT
			id,
			category_id,
			name,
			brand,
			model,
			license_plate,
			year,
			price_per_day,
			status,
			created_at
		FROM vehicles
		WHERE id = $1
	`

	var vehicle model.Vehicle

	err := r.db.QueryRow(ctx, query, id).Scan(
		&vehicle.ID,
		&vehicle.CategoryID,
		&vehicle.Name,
		&vehicle.Brand,
		&vehicle.Model,
		&vehicle.LicensePlate,
		&vehicle.Year,
		&vehicle.PricePerDay,
		&vehicle.Status,
		&vehicle.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("kendaraan dengan id %d tidak ditemukan", id)
		}

		return nil, fmt.Errorf("gagal mengambil kendaraan: %w", err)
	}

	return &vehicle, nil
}

func (r *VehicleRepository) Update(ctx context.Context, vehicle *model.Vehicle) error {
	query := `
		UPDATE vehicles
		SET
			category_id = $1,
			name = $2,
			brand = $3,
			model = $4,
			license_plate = $5,
			year = $6,
			price_per_day = $7,
			status = $8
		WHERE id = $9
	`

	result, err := r.db.Exec(
		ctx,
		query,
		vehicle.CategoryID,
		vehicle.Name,
		vehicle.Brand,
		vehicle.Model,
		vehicle.LicensePlate,
		vehicle.Year,
		vehicle.PricePerDay,
		vehicle.Status,
		vehicle.ID,
	)

	if err != nil {
		return fmt.Errorf("gagal mengubah kendaraan: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("kendaraan dengan id %d tidak ditemukan", vehicle.ID)
	}

	return nil
}

func (r *VehicleRepository) Delete(ctx context.Context, id int) error {
	query := `
		DELETE FROM vehicles
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus kendaraan: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("kendaraan dengan id %d tidak ditemukan", id)
	}

	return nil
}