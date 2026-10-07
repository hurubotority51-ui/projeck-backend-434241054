package service

import (
	"context"
	"fmt"
	"time"

	"projek-backend/app/model"
	"projek-backend/app/repository"
)

type RentalService struct {
	rentalRepository  *repository.RentalRepository
	vehicleRepository *repository.VehicleRepository
}

func NewRentalService(
	rentalRepository *repository.RentalRepository,
	vehicleRepository *repository.VehicleRepository,
) *RentalService {
	return &RentalService{
		rentalRepository:  rentalRepository,
		vehicleRepository: vehicleRepository,
	}
}

func (s *RentalService) Create(
	ctx context.Context,
	userID int,
	request *model.CreateRentalRequest,
) (*model.Rental, error) {

	if userID <= 0 {
		return nil, fmt.Errorf("user tidak valid")
	}

	vehicle, err := s.vehicleRepository.FindByID(
		ctx,
		request.VehicleID,
	)
	if err != nil {
		return nil, err
	}

	if vehicle.Status != "available" {
		return nil, fmt.Errorf("kendaraan tidak tersedia")
	}

	startDate, err := time.Parse(
		"2006-01-02",
		request.StartDate,
	)
	if err != nil {
		return nil, fmt.Errorf("format start_date harus YYYY-MM-DD")
	}

	endDate, err := time.Parse(
		"2006-01-02",
		request.EndDate,
	)
	if err != nil {
		return nil, fmt.Errorf("format end_date harus YYYY-MM-DD")
	}

	if endDate.Before(startDate) {
		return nil, fmt.Errorf("end_date tidak boleh sebelum start_date")
	}

	totalDays := int(endDate.Sub(startDate).Hours()/24) + 1

	if totalDays <= 0 {
		return nil, fmt.Errorf("total hari rental tidak valid")
	}

	totalPrice := float64(totalDays) * vehicle.PricePerDay

	rental := &model.Rental{
		UserID:      userID,
		VehicleID:   vehicle.ID,
		StartDate:   startDate,
		EndDate:     endDate,
		TotalDays:   totalDays,
		PricePerDay: vehicle.PricePerDay,
		TotalPrice:  totalPrice,
		Status:      "pending",
	}

	if err := s.rentalRepository.Create(
		ctx,
		rental,
	); err != nil {
		return nil, err
	}

	return rental, nil
}

func (s *RentalService) GetByID(
	ctx context.Context,
	id int,
	userID int,
	role string,
) (*model.Rental, error) {

	if id <= 0 {
		return nil, fmt.Errorf("id rental tidak valid")
	}

	rental, err := s.rentalRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if role == "customer" && rental.UserID != userID {
		return nil, fmt.Errorf("akses rental ditolak")
	}

	if role != "customer" && role != "staff" && role != "admin" {
		return nil, fmt.Errorf("role tidak valid")
	}

	return rental, nil
}

func (s *RentalService) FindAll(
	ctx context.Context,
	userID int,
	role string,
) ([]model.Rental, error) {

	if userID <= 0 {
		return nil, fmt.Errorf("user tidak valid")
	}

	if role != "customer" && role != "staff" && role != "admin" {
		return nil, fmt.Errorf("role tidak valid")
	}

	rentals, err := s.rentalRepository.FindAll(
		ctx,
		userID,
		role,
	)
	if err != nil {
		return nil, err
	}

	return rentals, nil
}
