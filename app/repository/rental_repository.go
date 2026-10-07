package repository

import (
	"context"
	"fmt"

	"projek-backend/app/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RentalRepository struct {
	db *pgxpool.Pool
}

func NewRentalRepository(db *pgxpool.Pool) *RentalRepository {
	return &RentalRepository{
		db: db,
	}
}

func (r *RentalRepository) Create(
	ctx context.Context,
	rental *model.Rental,
) error {

	query := `
		INSERT INTO rentals (
			user_id,
			vehicle_id,
			start_date,
			end_date,
			total_days,
			price_per_day,
			total_price,
			status
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		rental.UserID,
		rental.VehicleID,
		rental.StartDate,
		rental.EndDate,
		rental.TotalDays,
		rental.PricePerDay,
		rental.TotalPrice,
		rental.Status,
	).Scan(
		&rental.ID,
		&rental.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("gagal membuat rental: %w", err)
	}

	return nil
}

func (r *RentalRepository) FindByID(
	ctx context.Context,
	id int,
) (*model.Rental, error) {

	query := `
		SELECT
			id,
			user_id,
			vehicle_id,
			start_date,
			end_date,
			total_days,
			price_per_day,
			total_price,
			status,
			created_at
		FROM rentals
		WHERE id = $1
	`

	var rental model.Rental

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&rental.ID,
		&rental.UserID,
		&rental.VehicleID,
		&rental.StartDate,
		&rental.EndDate,
		&rental.TotalDays,
		&rental.PricePerDay,
		&rental.TotalPrice,
		&rental.Status,
		&rental.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("rental dengan id %d tidak ditemukan", id)
		}

		return nil, fmt.Errorf("gagal mencari rental: %w", err)
	}

	return &rental, nil
}

func (r *RentalRepository) FindAll(
	ctx context.Context,
	userID int,
	role string,
) ([]model.Rental, error) {

	var rows pgx.Rows
	var err error

	if role == "customer" {
		query := `
			SELECT
				id,
				user_id,
				vehicle_id,
				start_date,
				end_date,
				total_days,
				price_per_day,
				total_price,
				status,
				created_at
			FROM rentals
			WHERE user_id = $1
			ORDER BY id DESC
		`

		rows, err = r.db.Query(
			ctx,
			query,
			userID,
		)
	} else {
		query := `
			SELECT
				id,
				user_id,
				vehicle_id,
				start_date,
				end_date,
				total_days,
				price_per_day,
				total_price,
				status,
				created_at
			FROM rentals
			ORDER BY id DESC
		`

		rows, err = r.db.Query(
			ctx,
			query,
		)
	}

	if err != nil {
		return nil, fmt.Errorf(
			"gagal mengambil daftar rental: %w",
			err,
		)
	}

	defer rows.Close()

	var rentals []model.Rental

	for rows.Next() {
		var rental model.Rental

		err := rows.Scan(
			&rental.ID,
			&rental.UserID,
			&rental.VehicleID,
			&rental.StartDate,
			&rental.EndDate,
			&rental.TotalDays,
			&rental.PricePerDay,
			&rental.TotalPrice,
			&rental.Status,
			&rental.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf(
				"gagal membaca data rental: %w",
				err,
			)
		}

		rentals = append(rentals, rental)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"gagal membaca data rental: %w",
			err,
		)
	}

	return rentals, nil
}
