package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"projek-backend/app/handler"
	"projek-backend/app/repository"
	"projek-backend/app/service"
	"projek-backend/config"
	"projek-backend/middleware"
	"projek-backend/route"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	api := app.Group("/api/v1")

	api.Get("/protected", middleware.AuthMiddleware(), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "akses berhasil",
		})
	})

	api.Get(
		"/admin-test",
		middleware.AuthMiddleware(),
		middleware.RoleMiddleware("admin"),
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"success": true,
				"message": "akses admin berhasil",
			})
		},
	)

	userrepository := repository.NewUserRepository(config.DB)

	authservice := service.NewAuthService(userrepository)

	authhandler := handler.NewAuthHandler(authservice)

	route.AuthRoutes(api, authhandler)

	/////vehicle
	vehclerepository := repository.NewVehicleRepository(config.DB)

	vehicleservice := service.NewVehicleService(vehclerepository)

	vehiclehandler := handler.NewVehicleHandler(vehicleservice)

	route.VehicleRoutes(api, vehiclehandler)

	/////category
	categoryrepository := repository.NewVehicleCategoryRepository(config.DB)

	categoryservice := service.NewVehicleCategoryService(categoryrepository)

	categoryhandler := handler.NewVehicleCategoryHandler(categoryservice)

	route.VehicleCategoryRoutes(api, categoryhandler)
	//////rantal
	rentalrepository := repository.NewRentalRepository(config.DB)

	rentalservice := service.NewRentalService(
		rentalrepository,
		vehclerepository,
	)
	rentalhandler := handler.NewRentalHandler(rentalservice)
	route.RentalRoutes(api, rentalhandler)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Vehicle Rental API berjalan",
			"success": true,
		})
	})
	paymentrepository := repository.NewPaymentRepository(config.DB)

	paymentservice := service.NewPaymentService(
		paymentrepository,
		rentalrepository,
	)

	paymenthandler := handler.NewPaymentHandler(paymentservice)

	route.PaymentRoutes(api, paymenthandler)
	log.Fatal(app.Listen(":3000"))
}
