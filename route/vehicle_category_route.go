package route

import (
	"projek-backend/app/handler"

	"github.com/gofiber/fiber/v2"
)

func VehicleCategoryRoutes(api fiber.Router, categoryHandler *handler.VehicleCategoryHandler) {
	categories := api.Group("/categories")

	categories.Post("/", categoryHandler.Create)
	categories.Get("/", categoryHandler.GetAll)
	categories.Get("/:id", categoryHandler.GetByID)
	categories.Put("/:id", categoryHandler.Update)
	categories.Delete("/:id", categoryHandler.Delete)
}