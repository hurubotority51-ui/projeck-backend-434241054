package route

import (
	"projek-backend/app/handler"
	"projek-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func VehicleRoutes(api fiber.Router, vehicleHandler *handler.VehicleHandler) {
	vehicles := api.Group("/vehicles")

	vehicles.Post(
		"/",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("admin", "staff"),
		vehicleHandler.Create,
	)

	vehicles.Get("/", vehicleHandler.GetAll)
	vehicles.Get("/:id", vehicleHandler.GetByID)

	vehicles.Put(
		"/:id",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("admin", "staff"),
		vehicleHandler.Update,
	)

	vehicles.Delete(
		"/:id",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("admin"),
		vehicleHandler.Delete,
	)
}