package route

import (
	"projek-backend/app/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router, authHandler *handler.AuthHandler) {
	auth := api.Group("/auth")

	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
}