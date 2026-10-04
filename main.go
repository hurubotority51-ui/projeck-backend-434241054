package main

import (
	"log"

	"projek-backend/app/handler"
	"projek-backend/app/repository"
	"projek-backend/app/service"
	"projek-backend/config"
	"projek-backend/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	api := app.Group("/api/v1")

	vehclerepository := repository.NewVehicleRepository(config.DB)

	vehicleservice := service.NewVehicleService(vehclerepository)

	vehiclehandler := handler.NewVehicleHandler(vehicleservice)

	route.VehicleRoutes(api, vehiclehandler)


	categoryrepository := repository.NewVehicleCategoryRepository(config.DB)

	categoryservice := service.NewVehicleCategoryService(categoryrepository)

	categoryhandler := handler.NewVehicleCategoryHandler(categoryservice)

	route.VehicleCategoryRoutes(api, categoryhandler)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Vehicle Rental API berjalan",
			"success": true,
		})
	})

	log.Fatal(app.Listen(":3000"))
}