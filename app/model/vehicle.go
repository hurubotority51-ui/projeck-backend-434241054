package model
import "time"

type Vehicle struct {
	ID           int       `json:"id"`
	CategoryID   int       `json:"category_id"`
	Name         string    `json:"name"`
	Brand        string    `json:"brand"`
	Model        string    `json:"model"`
	LicensePlate string    `json:"license_plate"`
	Year         int       `json:"year"`
	PricePerDay  float64   `json:"price_per_day"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
}