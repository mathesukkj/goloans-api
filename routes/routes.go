package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/mathesukkj/goloans-api/controllers"
)

func NewRouter() *fiber.App {
	app := fiber.New()

	app.Post("/customer-loans", controllers.GetPossibleCustomerLoans)

	return app
}
