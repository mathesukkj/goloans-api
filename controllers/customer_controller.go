package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"

	"github.com/mathesukkj/goloans-api/models"
	"github.com/mathesukkj/goloans-api/services"
)

func CheckAvailableCustomerLoans(c fiber.Ctx) error {
	validate := validator.New()
	customer := new(models.Customer)

	if err := c.Bind().JSON(customer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error during body binding",
			"error":   err,
		})
	}

	if err := validate.Struct(customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error while parsing body. please check the correct format in the docs",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(models.CustomerLoansResponse{
		Customer: customer.Name,
		Loans:    services.GetAvailableLoans(customer),
	})
}
