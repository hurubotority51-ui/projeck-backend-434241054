package handler
import (
	"strconv"

	"projek-backend/app/model"
	"projek-backend/app/service"

	"github.com/gofiber/fiber/v2"
)

type VehicleHandler struct {
	service *service.VehicleService
}

func NewVehicleHandler(service *service.VehicleService) *VehicleHandler {
	return &VehicleHandler{
		service: service,
	}
}

func (h *VehicleHandler) Create(c *fiber.Ctx) error {
	var vehicle model.Vehicle

	if err := c.BodyParser(&vehicle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "format request tidak valid",
		})
	}

	if err := h.service.Create(c.Context(), &vehicle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "kendaraan berhasil ditambahkan",
		"data":    vehicle,
	})
}

func (h *VehicleHandler) GetAll(c *fiber.Ctx) error {
	vehicles, err := h.service.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "daftar kendaraan berhasil diambil",
		"data":    vehicles,
	})
}

func (h *VehicleHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "id kendaraan tidak valid",
		})
	}

	vehicle, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "detail kendaraan berhasil diambil",
		"data":    vehicle,
	})
}

func (h *VehicleHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "id kendaraan tidak valid",
		})
	}

	var vehicle model.Vehicle

	if err := c.BodyParser(&vehicle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "format request tidak valid",
		})
	}

	vehicle.ID = id

	if err := h.service.Update(c.Context(), &vehicle); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "kendaraan berhasil diperbarui",
		"data":    vehicle,
	})
}

func (h *VehicleHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "id kendaraan tidak valid",
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
		"message": "kendaraan berhasil dihapus",
	})
}