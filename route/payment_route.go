package route

import (
	"projek-backend/app/handler"
	"projek-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func PaymentRoutes(
	api fiber.Router,
	paymentHandler *handler.PaymentHandler,
) {
	payments := api.Group("/payments")

	payments.Post(
		"/",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("customer", "staff", "admin"),
		paymentHandler.Create,
	)

	payments.Get(
		"/:id",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("customer", "staff", "admin"),
		paymentHandler.GetByID,
	)
}
