package route

import (
	"projek-backend/app/handler"
	"projek-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func RentalRoutes(
	api fiber.Router,
	rentalHandler *handler.RentalHandler,
) {
	rentals := api.Group("/rentals")

	rentals.Post(
		"/",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("customer", "staff", "admin"),
		rentalHandler.Create,
	)

	rentals.Get(
		"/",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("customer", "staff", "admin"),
		rentalHandler.GetAll,
	)

	rentals.Get(
		"/:id",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("customer", "staff", "admin"),
		rentalHandler.GetByID,
	)
}
