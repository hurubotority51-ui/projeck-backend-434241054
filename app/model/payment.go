package model

import "time"

type Payment struct {
	ID            int        `json:"id"`
	RentalID      int        `json:"rental_id"`
	Amount        float64    `json:"amount"`
	PaymentMethod string     `json:"payment_method"`
	PaymentStatus string     `json:"payment_status"`
	PaidAt        *time.Time `json:"paid_at"`
	CreatedAt     time.Time  `json:"created_at"`
}

type CreatePaymentRequest struct {
	RentalID      int     `json:"rental_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
}
