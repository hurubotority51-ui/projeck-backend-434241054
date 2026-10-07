package handler

import (
	"strconv"

	"projek-backend/app/model"
	"projek-backend/app/service"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type RentalHandler struct {
	service *service.RentalService
}

func NewRentalHandler(service *service.RentalService) *RentalHandler {
	return &RentalHandler{
		service: service,
	}
}

func (h *RentalHandler) Create(c *fiber.Ctx) error {
	var request model.CreateRentalRequest

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "format request tidak valid",
		})
	}

	claims, ok := c.Locals("claims").(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "data user tidak ditemukan",
		})
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "user id tidak valid",
		})
	}

	userID := int(userIDFloat)

	rental, err := h.service.Create(
		c.Context(),
		userID,
		&request,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "rental berhasil dibuat",
		"data":    rental,
	})
}

func (h *RentalHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "id rental tidak valid",
		})
	}

	claims, ok := c.Locals("claims").(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "data user tidak ditemukan",
		})
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "user id tidak valid",
		})
	}

	userID := int(userIDFloat)

	role, ok := claims["role"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "role user tidak ditemukan",
		})
	}

	rental, err := h.service.GetByID(
		c.Context(),
		id,
		userID,
		role,
	)

	if err != nil {
		if err.Error() == "akses rental ditolak" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "data rental berhasil diambil",
		"data":    rental,
	})
}

func (h *RentalHandler) GetAll(c *fiber.Ctx) error {
	claims, ok := c.Locals("claims").(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "data user tidak ditemukan",
		})
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "user id tidak valid",
		})
	}

	userID := int(userIDFloat)

	role, ok := claims["role"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "role user tidak ditemukan",
		})
	}

	rentals, err := h.service.FindAll(
		c.Context(),
		userID,
		role,
	)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "daftar rental berhasil diambil",
		"data":    rentals,
	})
}
