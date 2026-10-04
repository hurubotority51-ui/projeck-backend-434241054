package main

import (
	"log"

	"projek-backend/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Vehicle Rental API berjalan",
			"success": true,
		})
	})

	log.Fatal(app.Listen(":3000"))
}