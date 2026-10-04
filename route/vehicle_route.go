package route
import (
	"projek-backend/app/handler"

	"github.com/gofiber/fiber/v2"
)

func VehicleRoutes(api fiber.Router, vehicleHandler *handler.VehicleHandler) {
	vehicles := api.Group("/vehicles")

	vehicles.Post("/", vehicleHandler.Create)
	vehicles.Get("/", vehicleHandler.GetAll)
	vehicles.Get("/:id", vehicleHandler.GetByID)
	vehicles.Put("/:id", vehicleHandler.Update)
	vehicles.Delete("/:id", vehicleHandler.Delete)
}