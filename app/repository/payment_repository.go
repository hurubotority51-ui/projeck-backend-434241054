package repository

import (
	"context"
	"fmt"

	"projek-backend/app/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepository struct {
	db *pgxpool.Pool
}

func NewPaymentRepository(db *pgxpool.Pool) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}

func (r *PaymentRepository) Create(
	ctx context.Context,
	payment *model.Payment,
) error {

	query := `
		INSERT INTO payments (
			rental_id,
			amount,
			payment_method,
			payment_status,
			paid_at
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	err := r.db.QueryRow(
		ctx,
		query,
		payment.RentalID,
		payment.Amount,
		payment.PaymentMethod,
		payment.PaymentStatus,
		payment.PaidAt,
	).Scan(
		&payment.ID,
		&payment.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("gagal membuat payment: %w", err)
	}

	return nil
}

func (r *PaymentRepository) FindByID(
	ctx context.Context,
	id int,
) (*model.Payment, error) {

	query := `
		SELECT
			id,
			rental_id,
			amount,
			payment_method,
			payment_status,
			paid_at,
			created_at
		FROM payments
		WHERE id = $1
	`

	var payment model.Payment

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&payment.ID,
		&payment.RentalID,
		&payment.Amount,
		&payment.PaymentMethod,
		&payment.PaymentStatus,
		&payment.PaidAt,
		&payment.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf(
				"payment dengan id %d tidak ditemukan",
				id,
			)
		}

		return nil, fmt.Errorf(
			"gagal mencari payment: %w",
			err,
		)
	}

	return &payment, nil
}

func (r *PaymentRepository) FindByRentalID(
	ctx context.Context,
	rentalID int,
) (*model.Payment, error) {

	query := `
		SELECT
			id,
			rental_id,
			amount,
			payment_method,
			payment_status,
			paid_at,
			created_at
		FROM payments
		WHERE rental_id = $1
		ORDER BY id DESC
		LIMIT 1
	`

	var payment model.Payment

	err := r.db.QueryRow(
		ctx,
		query,
		rentalID,
	).Scan(
		&payment.ID,
		&payment.RentalID,
		&payment.Amount,
		&payment.PaymentMethod,
		&payment.PaymentStatus,
		&payment.PaidAt,
		&payment.CreatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf(
				"payment untuk rental dengan id %d tidak ditemukan",
				rentalID,
			)
		}

		return nil, fmt.Errorf(
			"gagal mencari payment berdasarkan rental: %w",
			err,
		)
	}

	return &payment, nil
}
