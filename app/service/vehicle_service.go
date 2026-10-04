package service
import (
	"context"
	"fmt"
	"strings"

	"projek-backend/app/model"
	"projek-backend/app/repository"
)

type VehicleService struct {
	repository *repository.VehicleRepository
}

func NewVehicleService(repository *repository.VehicleRepository) *VehicleService {
	return &VehicleService{
		repository: repository,
	}
}

func (s *VehicleService) Create(ctx context.Context, vehicle *model.Vehicle) error {
	if strings.TrimSpace(vehicle.Name) == "" {
		return fmt.Errorf("nama kendaraan wajib diisi")
	}

	if strings.TrimSpace(vehicle.Brand) == "" {
		return fmt.Errorf("merek kendaraan wajib diisi")
	}

	if strings.TrimSpace(vehicle.LicensePlate) == "" {
		return fmt.Errorf("nomor plat wajib diisi")
	}

	if vehicle.CategoryID <= 0 {
		return fmt.Errorf("category_id tidak valid")
	}

	if vehicle.PricePerDay <= 0 {
		return fmt.Errorf("harga sewa per hari harus lebih dari 0")
	}

	if vehicle.Year <= 0 {
		return fmt.Errorf("tahun kendaraan tidak valid")
	}

	if vehicle.Status == "" {
		vehicle.Status = "available"
	}

	return s.repository.Create(ctx, vehicle)
}

func (s *VehicleService) GetAll(ctx context.Context) ([]model.Vehicle, error) {
	return s.repository.FindAll(ctx)
}

func (s *VehicleService) GetByID(ctx context.Context, id int) (*model.Vehicle, error) {
	if id <= 0 {
		return nil, fmt.Errorf("id kendaraan tidak valid")
	}

	return s.repository.FindByID(ctx, id)
}

func (s *VehicleService) Update(ctx context.Context, vehicle *model.Vehicle) error {
	if vehicle.ID <= 0 {
		return fmt.Errorf("id kendaraan tidak valid")
	}

	if strings.TrimSpace(vehicle.Name) == "" {
		return fmt.Errorf("nama kendaraan wajib diisi")
	}

	if strings.TrimSpace(vehicle.Brand) == "" {
		return fmt.Errorf("merek kendaraan wajib diisi")
	}

	if strings.TrimSpace(vehicle.LicensePlate) == "" {
		return fmt.Errorf("nomor plat wajib diisi")
	}

	if vehicle.CategoryID <= 0 {
		return fmt.Errorf("category_id tidak valid")
	}

	if vehicle.PricePerDay <= 0 {
		return fmt.Errorf("harga sewa per hari harus lebih dari 0")
	}

	if vehicle.Year <= 0 {
		return fmt.Errorf("tahun kendaraan tidak valid")
	}

	if vehicle.Status == "" {
		vehicle.Status = "available"
	}

	return s.repository.Update(ctx, vehicle)
}

func (s *VehicleService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("id kendaraan tidak valid")
	}

	return s.repository.Delete(ctx, id)
}