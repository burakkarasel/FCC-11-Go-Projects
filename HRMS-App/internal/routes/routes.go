package routes

import (
	"github.com/burakkarasel/HRMS-App/internal/api"
	"github.com/gofiber/fiber/v2"
)

// SetUpRoutes sets up the routes for the fiber app
func SetUpRoutes(app *fiber.App) {
	app.Get("/employee", api.ListEmployees)
	app.Get("/employee/:id", api.GetEmployee)
	app.Post("/employee", api.NewEmployee)
	app.Put("/employee/:id", api.UpdateEmployee)
	app.Delete("/employee/:id", api.DeleteEmployee)
}
