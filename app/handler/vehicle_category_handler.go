package handler

import (
	"strconv"

	"projek-backend/app/model"
	"projek-backend/app/service"

	"github.com/gofiber/fiber/v2"
)

type VehicleCategoryHandler struct {
	service *service.VehicleCategoryService
}

func NewVehicleCategoryHandler(service *service.VehicleCategoryService) *VehicleCategoryHandler {
	return &VehicleCategoryHandler{
		service: service,
	}
}

func (h *VehicleCategoryHandler) Create(c *fiber.Ctx) error {
	var category model.VehicleCategory

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "format request tidak valid",
		})
	}

	if err := h.service.Create(c.Context(), &category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "kategori kendaraan berhasil ditambahkan",
		"data":    category,
	})
}

func (h *VehicleCategoryHandler) GetAll(c *fiber.Ctx) error {
	categories, err := h.service.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "daftar kategori kendaraan berhasil diambil",
		"data":    categories,
	})
}

func (h *VehicleCategoryHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "id kategori tidak valid",
		})
	}

	category, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "detail kategori kendaraan berhasil diambil",
		"data":    category,
	})
}

func (h *VehicleCategoryHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "id kategori tidak valid",
		})
	}

	var category model.VehicleCategory

	if err := c.BodyParser(&category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "format request tidak valid",
		})
	}

	category.ID = id

	if err := h.service.Update(c.Context(), &category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "kategori kendaraan berhasil diperbarui",
		"data":    category,
	})
}

func (h *VehicleCategoryHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "id kategori tidak valid",
		})
	}

	if err := h.service.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "kategori kendaraan berhasil dihapus",
	})
}