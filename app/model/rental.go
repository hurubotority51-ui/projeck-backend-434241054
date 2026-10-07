package model

import "time"

type Rental struct {
	ID           int       `json:"id"`
	UserID       int       `json:"user_id"`
	VehicleID    int       `json:"vehicle_id"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	TotalDays    int       `json:"total_days"`
	PricePerDay  float64   `json:"price_per_day"`
	TotalPrice   float64   `json:"total_price"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}

type CreateRentalRequest struct {
	VehicleID int    `json:"vehicle_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}