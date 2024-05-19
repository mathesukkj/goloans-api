package controllers

import "github.com/gofiber/fiber/v3"

func CustomerLoans(c fiber.Ctx) error {
	return c.SendString("hello world!")
}
