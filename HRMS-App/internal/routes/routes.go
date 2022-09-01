package routes

import (
	"github.com/burakkarasel/HRMS-App/internal/api"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App) {
	app.Get("/employee", api.GetEmployees)
	app.Post("/employee", api.NewEmployee)
	app.Put("/employee/:id", api.UpdateEmployee)
	app.Delete("/employee/:id", api.DeleteEmployee)
}
