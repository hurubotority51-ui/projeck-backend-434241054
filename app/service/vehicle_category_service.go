package service

import (
	"context"
	"fmt"
	"strings"

	"projek-backend/app/model"
	"projek-backend/app/repository"
)

type VehicleCategoryService struct {
	repository *repository.VehicleCategoryRepository
}

func NewVehicleCategoryService(repository *repository.VehicleCategoryRepository) *VehicleCategoryService {
	return &VehicleCategoryService{
		repository: repository,
	}
}

func (s *VehicleCategoryService) Create(ctx context.Context, category *model.VehicleCategory) error {
	if strings.TrimSpace(category.Name) == "" {
		return fmt.Errorf("nama kategori wajib diisi")
	}

	if len(strings.TrimSpace(category.Name)) > 50 {
		return fmt.Errorf("nama kategori maksimal 50 karakter")
	}

	return s.repository.Create(ctx, category)
}

func (s *VehicleCategoryService) GetAll(ctx context.Context) ([]model.VehicleCategory, error) {
	return s.repository.FindAll(ctx)
}

func (s *VehicleCategoryService) GetByID(ctx context.Context, id int) (*model.VehicleCategory, error) {
	if id <= 0 {
		return nil, fmt.Errorf("id kategori tidak valid")
	}

	return s.repository.FindByID(ctx, id)
}

func (s *VehicleCategoryService) Update(ctx context.Context, category *model.VehicleCategory) error {
	if category.ID <= 0 {
		return fmt.Errorf("id kategori tidak valid")
	}

	if strings.TrimSpace(category.Name) == "" {
		return fmt.Errorf("nama kategori wajib diisi")
	}

	if len(strings.TrimSpace(category.Name)) > 50 {
		return fmt.Errorf("nama kategori maksimal 50 karakter")
	}

	return s.repository.Update(ctx, category)
}

func (s *VehicleCategoryService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return fmt.Errorf("id kategori tidak valid")
	}

	return s.repository.Delete(ctx, id)
}